// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package main

import (
	"crypto/md5"
	"fmt"
	"net/http"

	jira "github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
)

type jiraTestInstance struct {
	JIRAInstance
}

var _ Instance = (*jiraTestInstance)(nil)

const (
	mockCurrentInstanceURL = "http://jiraTestInstanceURL.some"
)

func keyWithMockInstance(key string) string {
	h := md5.New()
	fmt.Fprintf(h, "%s/%s", mockCurrentInstanceURL, key)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (jti jiraTestInstance) GetURL() string {
	return mockCurrentInstanceURL
}
func (jti jiraTestInstance) GetMattermostKey() string {
	return "jiraTestInstanceMattermostKey"
}
func (jti jiraTestInstance) GetDisplayDetails() map[string]string {
	return map[string]string{}
}
func (jti jiraTestInstance) GetUserConnectURL(mattermostUserId string) (string, error) {
	return "http://jiraTestInstanceUserConnectURL.some", nil
}

func (jti jiraTestInstance) GetClient(jiraUser JIRAUser) (Client, error) {
	client, err := jira.NewClient(&http.Client{}, "testClient")
	if err != nil {
		return nil, err
	}
	return newCloudClient(client), nil
}

func (jti jiraTestInstance) GetUserGroups(jiraUser JIRAUser) ([]*jira.UserGroup, error) {
	return nil, errors.New("not implemented")
}

type mockCurrentInstanceStore struct {
	plugin *Plugin
}

func (store mockCurrentInstanceStore) StoreCurrentJIRAInstance(ji Instance) error {
	return nil
}
func (store mockCurrentInstanceStore) LoadCurrentJIRAInstance() (Instance, error) {
	return &jiraTestInstance{
		JIRAInstance: *NewJIRAInstance(store.plugin, "test", "jiraTestInstanceKey"),
	}, nil
}

type mockCurrentInstanceStoreNoInstance struct {
	plugin *Plugin
}

func (store mockCurrentInstanceStoreNoInstance) StoreCurrentJIRAInstance(ji Instance) error {
	return nil
}
func (store mockCurrentInstanceStoreNoInstance) LoadCurrentJIRAInstance() (Instance, error) {
	return nil, errors.New("failed to load current Jira instance: not found")
}

type mockUserStore struct {
	loadJiraUserResp JIRAUser
}

func (store mockUserStore) StoreUserInfo(ji Instance, mattermostUserId string, jiraUser JIRAUser) error {
	return nil
}
func (store mockUserStore) LoadJIRAUser(ji Instance, mattermostUserId string) (JIRAUser, error) {
	return store.loadJiraUserResp, nil
}
func (store mockUserStore) LoadMattermostUserId(ji Instance, jiraUserName string) (string, error) {
	return "testMattermostUserId012345", nil
}
func (store mockUserStore) LoadJIRAUserByAccountId(ji Instance, accoundId string) (JIRAUser, error) {
	return JIRAUser{}, nil
}
func (store mockUserStore) DeleteUserInfo(ji Instance, mattermostUserId string) error {
	return nil
}
