**Mattermost Preview**  
**Team Communication Service**  
**Version SUPER AWESOME**


About Mattermost
================

Mattermost is a team communication service. It brings team messaging and file sharing into one place, accessible across PCs and phones, with archiving and search.

We built Mattermost to help teams focus on what matters most to them. It works for us, we hope it works for you too.

Learn More
==========
<ul>
<li/>Ask the core team anything at: http://forum.mattermost.org</li>
<li/>Share feature requests and upvotes: http://www.mattermost.org/feature-requests/</li>
<li/>File bugs: http://www.mattermost.org/filing-issues/</li>
<li/>Make a pull request: http://www.mattermost.org/contribute-to-mattermost/</li>
</ul>

Installing the Mattermost
=========================

You're installing "Mattermost Preview", a pre-released 0.5.0 version intended for an early look at what we're building. While SpinPunch runs this version internally, it's not recommended for production deployments since we can't guarantee API stability or backwards compatibility until our 1.0.0 version release. 

That said, any issues at all, please let us know on the Mattermost forum at: http://forum.mattermost.org 

Local Machine Setup (Docker)
-----------------------------

### Mac OSX ###

1. Follow the instructions at http://docs.docker.com/installation/mac/  
    1. Use the Boot2Docker command-line utility  
    2. If you do command-line setup use: `boot2docker init eval “$(boot2docker shellinit)”`  
2. Get your Docker IP address with `boot2docker ip`
3. Add a line to your /etc/hosts that goes `<Docker IP> dockerhost`
4. Run `boot2docker shellinit` and copy the export statements to your ~/.bash\_profile
5. Run `docker run --name mattermost-dev -d --publish 8065:80 mattermost/platform:helium`. 
6. When docker is done fetching the image, open http://dockerhost:8065/ in your browser

### Ubuntu ###
1. Follow the instructions at https://docs.docker.com/installation/ubuntulinux/ or use the summary below.

	``` bash
	sudo apt-get update
	sudo apt-get install wget
	wget -qO- https://get.docker.com/ | sh
	sudo usermod -aG docker <username>
	sudo service docker start
	newgrp docker
	```

2. Run `docker run --name mattermost-dev -d --publish 8065:80 mattermost/platform:helium`
3. When docker is done fetching the image, open http://localhost:8065/ in your browser

### Arch ###
1. Install docker using the following commands:

	``` bash
	pacman -S docker
	systemctl enable docker.service
	systemctl start docker.service
	gpasswd -a <username> docker
	newgrp docker
	```

2. Start docker container:

	``` bash
	docker run --name mattermost-dev -d --publish 8065:80 mattermost/platform:helium
	```

3. When docker is done fetching the image, open http://localhost:8065/ in your browser.

### Notes ###
If your ISP blocks port 25 then you may install locally but email will not be sent.

If you want to work with the latest bits in the repo you can run the cmd  
`docker run --name mattermost-dev -d --publish 8065:80 mattermost/platform:latest`

You can update to the latest bits by running  
`docker pull mattermost/platform:latest`

If you wish to remove mattermost-dev use the following commands  

1. `docker stop mattermost-dev`
2. `docker rm -v mattermost-dev`

If you wish to gain access to the container use the following commands
1. `docker exec -ti mattermost-dev /bin/bash`

AWS Elastic Beanstalk Setup (Docker)
------------------------------------

1. Create a new elastic beanstalk docker application using the Dockerrun.aws.json file provided. 
	1. From the AWS console select Elastic Beanstalk
	2. Select "Create New Application" from the top right.
	3. Name the application and press next
	4. Select "Create a web server" environment.
	5. If asked, select create and IAM role and instance profile and press next.
	6. For predefined configuration select docker. For environment type select single instance. 
	7. For application source, select upload your own and upload Dockerrun.aws.json from docker/Dockerrun.aws.json. Everything else may be left at default.
	8. Select an environment name, this is how you will refer to your environment. Make sure the URL is available then press next.
	9. The options on the additional resources page may be left at default unless you wish to change them. Press Next.
	10. On the configuration details place. Select an instance type of t2.small or larger.
	11. You can set the configuration details as you please but they may be left at their defaults. When you are done press next.
	12. Environment tags my be left blank. Press next.
	13. You will be asked to review your information. Press Launch.

4. Try it out!
	14. Wait for beanstalk to update the environment.
	15. Try it out by entering the domain of the form \*.elasticbeanstalk.com found at the top of the dashboard into your browser. You can also map your own domain if you wish. 

Contributing 
------------ 
 
To contribute to this open source project please review the Mattermost Contribution Guidelines at http://www.mattermost.org/contribute-to-mattermost/. 

License
-------

Mattermost is licensed under an "Apache-wrapped AGPL" model, which means you can run and link to the system using Configuration Files and Admin Tools licensed under Apache, version 2.0, as described in the LICENSE file, as an explicit exception to the terms of the GNU Affero General Public License (AGPL) that applies to most of the remaining source files. See individual files for details.

