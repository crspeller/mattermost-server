// Copyright (c) 2015 Spinpunch, Inc. All Rights Reserved.
// See License.txt for license information.




var utils = require('../utils/utils.jsx');
var client = require('../utils/client.jsx');
var UserStore = require('../stores/user_store.jsx');
var BrowserStore = require('../stores/browser_store.jsx');
var Constants = require('../utils/constants.jsx');


var FindTeamDomain = React.createClass({
    handleSubmit: function(e) {
        e.preventDefault();
        var state = { }

        var urlID = this.refs.urlID.getDOMNode().value.trim();
        if (!urlID) {
            state.server_error = "A domain is required"
            this.setState(state);
            return;
        }

        if (!BrowserStore.isLocalStorageSupported()) {
            state.server_error = "This service requires local storage to be enabled. Please enable it or exit private browsing.";
            this.setState(state);
            return;
        }

        state.server_error = "";
        this.setState(state);

        client.findTeamByURLId(urlId,
            function(data) {
                console.log(data);
                if (data) {
					window.location.href = req.getResponseHeader(Constants.TEAM_URL_HEADER) + '/channels/town-square';
                }
                else {
                    this.state.server_error = "We couldn't find your " + strings.Team + ".";
                    this.setState(this.state);
                }
            }.bind(this),
            function(err) {
                this.state.server_error = err.message;
                this.setState(this.state);
            }.bind(this)
        );
    },
    getInitialState: function() {
        return { };
    },
    render: function() {
        var server_error = this.state.server_error ? <label className="control-label">{this.state.server_error}</label> : null;

        return (
            <div className="signup-team__container">
                <div>
                    <span className="signup-team__name">{ config.SiteName }</span>
                    <br/>
                    <span className="signup-team__subdomain">Enter your {strings.Team}'s domain.</span>
                    <br/>
                    <br/>
                </div>
                <form onSubmit={this.handleSubmit}>
                    <div className={server_error ? 'form-group has-error' : 'form-group'}>
                        { server_error }
                        <input type="text" className="form-control" name="domain" ref="domain" placeholder="team domain" />
                    </div>
                    <div className="form-group">
                        <button type="submit" className="btn btn-primary">Continue</button>
                    </div>
                    <div>
                        <span>Don't remember your {strings.Team}'s domain? <a href="/find_team">Find it here</a></span>
                    </div>
                    <br/>
                    <br/>
                    <br/>
                    <br/>
                    <br/>
                    <br/>
                    <div>
                        <span>{"Want to create your own " + strings.Team + "?"} <a href={utils.getHomeLink()} className="signup-team-login">Sign up now</a></span>
                    </div>
                </form>
            </div>
        );
    }
});

module.exports = React.createClass({
    handleSubmit: function(e) {
        e.preventDefault();
        var state = { }

        var urlId = this.refs.urlID.getDOMNode().value.trim();
        if (!urlId) {
            state.server_error = "A domain is required"
            this.setState(state);
            return;
        }

        var email = this.refs.email.getDOMNode().value.trim();
        if (!email) {
            state.server_error = "An email is required"
            this.setState(state);
            return;
        }

        var password = this.refs.password.getDOMNode().value.trim();
        if (!password) {
            state.server_error = "A password is required"
            this.setState(state);
            return;
        }

        if (!BrowserStore.isLocalStorageSupported()) {
            state.server_error = "This service requires local storage to be enabled. Please enable it or exit private browsing.";
            this.setState(state);
            return;
        }

        state.server_error = "";
        this.setState(state);

        client.loginByEmail(urlId, email, password,
            function(data) {
                UserStore.setLastURLId(urlId);
                UserStore.setLastEmail(email);
                UserStore.setCurrentUser(data);

                var redirect = utils.getUrlParameter("redirect");
                if (redirect) {
                    window.location.pathname = decodeURI(redirect);
                } else {
					if (this.props.urlMode == Constants.URL_MODE_PATH) {
						window.location.pathname = '/' + urlId + '/channels/town-square';
					} else {
						window.location.pathname = '/channels/town-square';
					}
                }

            }.bind(this),
            function(err) {
                if (err.message == "Login failed because email address has not been verified") {
					window.location.href = '/verify?urlId=' + encodeURIComponent(urlId) + '&email=' + encodeURIComponent(email);
                    return;
                }
                state.server_error = err.message;
                this.valid = false;
                this.setState(state);
            }.bind(this)
        );
    },
    getInitialState: function() {
        return { };
    },
    render: function() {
        var server_error = this.state.server_error ? <label className="control-label">{this.state.server_error}</label> : null;
        var priorEmail = UserStore.getLastEmail() !== "undefined" ? UserStore.getLastEmail() : ""

        var emailParam = utils.getUrlParameter("email");
        if (emailParam) {
            priorEmail = decodeURIComponent(emailParam);
        }

        var subDomainClass = "form-control hidden";
		var teamName = "";
		var domainName = "";
		var urlID = "";
		if (this.props.urlMode == Constants.URL_MODE_DOMAIN) {
			var subDomain = utils.getSubDomain();
			if (utils.isTestDomain()) {
				subDomainClass = "form-control";
				subDomain = UserStore.getLastURLId();
			} else if (subDomain == "") {
				return (<FindTeamDomain />);
			}
			teamName = subDomain;
			urlID = subDomain;
			domainName = utils.getDomainWithOutSub();
		} else { // URL MODE Path
			teamName = window.location.pathname.split('/')[1];
			urlID = teamName;
		}

        return (
            <div className="signup-team__container">
                <div>
                    <span className="signup-team__name">{ teamName }</span>
                    <br/>
                    <span className="signup-team__subdomain">{ domainName }</span>
                    <br/>
                    <br/>
                </div>
                <form onSubmit={this.handleSubmit}>
                    <div className={server_error ? 'form-group has-error' : 'form-group'}>
                        { server_error }
                        <input type="text" className={subDomainClass} name="urlID" defaultValue={urlID} ref="urlID" placeholder="Domain" />
                    </div>
                    <div className={server_error ? 'form-group has-error' : 'form-group'}>
                        <input type="email" className="form-control" name="email" defaultValue={priorEmail}  ref="email" placeholder="Email" />
                    </div>
                    <div className={server_error ? 'form-group has-error' : 'form-group'}>
                        <input type="password" className="form-control" name="password" ref="password" placeholder="Password" />
                    </div>
                    <div className="form-group">
                        <button type="submit" className="btn btn-primary">Sign in</button>
                    </div>
                    <div className="form-group form-group--small">
                        <span><a href="/find_team">{"Find other " + strings.TeamPlural}</a></span>
                    </div>
                    <div className="form-group">
                        <a href="/reset_password">I forgot my password</a>
                    </div>
                    <div className="external-link">
                        <span>{"Want to create your own " + strings.Team + "?"} <a href={utils.getHomeLink()} className="signup-team-login">Sign up now</a></span>
                    </div>
                </form>
            </div>
        );
    }
});
