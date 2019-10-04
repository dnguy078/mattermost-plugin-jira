import React from 'react';

import {ChannelSubscription} from 'types/model';

import {SharedProps} from './shared_props';
import ConfirmModal from './confirm_modal'

type Props = SharedProps & {
    showEditChannelSubscription: (subscription: ChannelSubscription) => void;
    showCreateChannelSubscription: () => void;
};

type State = {
    error: string | null;
    showConfirmModal: boolean | false;
}

export default class SelectChannelSubscriptionInternal extends React.PureComponent<Props, State> {
    state = {
        error: null,
        showConfirmModal: false,
    };

    handleDeactivateCancel = () => {
        this.setState({showConfirmModal: false});
    }

    handleConfirm = (sub: ChannelSubscription) => {
        this.setState({showConfirmModal: false});
        this.deleteChannelSubscription(sub)
    }

    deleteChannelSubscription = (sub: ChannelSubscription): void => {
        this.props.deleteChannelSubscription(sub).then((res: {error?: {message: string}}) => {
            if (res.error) {
                this.setState({error: res.error.message});
            }
        });
    };

    handleDeleteChannelSubscription = (): void => {
          this.setState({showConfirmModal: true});
    };

    render(): React.ReactElement {
        const {channel} = this.props;
        const {error, showConfirmModal} = this.state;

        const headerText = `Jira Subscriptions in "${channel.name}"`;

        let errorDisplay = null;
        if (error) {
            errorDisplay = (
                <span className='error'>{error}</span>
            );
        }

        return (
            <div>
                <h1>{headerText}</h1>
                <button
                    className='btn btn-info'
                    onClick={this.props.showCreateChannelSubscription}
                >
                    {'Create Subscription'}
                </button>
                {errorDisplay}
                {this.props.channelSubscriptions.map((sub) => (
                    <div
                        key={sub.id}
                        className='select-channel-subscriptions-row'
                    >
                        <ConfirmModal
                            cancelButtonText={'Cancel'}
                            confirmButtonText={'Confirm'}
                            confirmButtonClass={'btn btn-danger'}
                            hideCancel={false}
                            message={'Delete Subscription "' + sub.id + '"?'}
                            onCancel={this.handleDeactivateCancel}
                            onConfirm={(): void => this.handleConfirm(sub)}
                            show={showConfirmModal}
                            title={'Subscription'}
                        />
                        <div className='channel-subscription-id-container'>
                            <span>{sub.id}</span>
                        </div>
                        <button
                            className='btn btn-info'
                            onClick={(): void => this.props.showEditChannelSubscription(sub)}
                        >
                            {'Edit'}
                        </button>
                        <button
                            className='btn btn-danger'
                            onClick={(): void => this.handleDeleteChannelSubscription()}
                        >
                            {'Delete'}
                        </button>
                    </div>
                ))}
            </div>
        );
    }
}
