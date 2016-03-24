// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

import ThreadView from './thread_view.jsx';
import React from 'react';

export default class Thread extends React.Component {
    render() {
        return (
            <ThreadView threadRootPost={this.props.params.postid}/>
        );
    }
}

Thread.propTypes = {
    params: React.PropTypes.object
};
