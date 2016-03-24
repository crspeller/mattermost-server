// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

import React from 'react';

import PostStore from 'stores/post_store.jsx';

export default class ThreadView extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        /*const posts = PostStore.getSelectedPostThread(this.props.threadRootPost);
        var postsArray = [];

        for (const id in posts) {
            if (posts.hasOwnProperty(id)) {
                const cpost = posts[id];
                if (cpost.root_id === this.props.threadRootPost.id) {
                    postsArray.push(cpost);
                }
            }
            }*/
        return (
            <p>{'Five sandwitches'}</p>
        );
    }
}

ThreadView.propTypes = {
    threadRootPost: React.PropTypes.string
};
