// Copyright (c) 2015 Spinpunch, Inc. All Rights Reserved.
// See License.txt for license information.

module.exports = React.createClass({
	propTypes: {
		title: React.PropTypes.string.isRequired,
		submit: React.PropTypes.func.isRequired,
		client_error: React.PropTypes.string,
		server_error: React.PropTypes.string
	},
	getInitialState: function() {
		return {
			client_error: this.props.client_error,
			server_error: this.props.server_error
		};
	},
	componentWillReceiveProps: function() {
		this.setState({
			client_error: this.props.client_error,
			server_error: this.props.server_error
		});
	},
	doFileSelect: function(e) {
		e.preventDefault();
		this.setState({
			client_error: "",
			server_error: ""
		});
	},
	doSubmit: function(e) {
		e.preventDefault();
		var inputnode = this.refs.uploadinput.getDOMNode();
		if (inputnode.files && inputnode.files[0])
		{
			this.props.submit(inputnode.files[0]);
		}
		else
		{
			this.setState({client_error: "No file selected."});
		}
	},
	doCancel: function(e) {
		e.preventDefault();
		this.refs.uploadinput.getDOMNode().value = "";
		this.setState({
			client_error: "",
			server_error: ""
		});
	},
	render: function() {
        var client_error = this.state.client_error ? <div className='form-group has-error'><label className='control-label'>{ this.state.client_error }</label></div> : null;
        var server_error = this.state.server_error ? <div className='form-group has-error'><label className='control-label'>{ this.state.server_error }</label></div> : null;
		return (
			<ul className="section-max">
			<li className="col-xs-12 section-title">{this.props.title}</li>
			<li className="col-xs-offset-3 col-xs-8">
				<ul className="setting-list">
					<li className="setting-list-item">
						{ server_error }
						{ client_error }
						<span className="btn btn-sm btn-primary btn-file sel-btn">SelectFile<input ref="uploadinput" accept=".zip" type="file" onchange={this.onFileSelect}/></span>
						<a className={"btn btn-sm btn-primary"} onClick={this.doSubmit}>Import</a>
						<a className="btn btn-sm theme" href="#" onClick={this.doCancel}>Cancel</a>
					</li>
				</ul>
			</li>
		</ul>
		);
	}
});
