#BASE IMAGE
FROM jenkins/jenkins:latest

# The author
LABEL author="Author Name"

#Setting up environment variables

#Disabling the Setup Wizard
ENV JAVA_OPTS -Djenkins.install.runSetupWizard=false
#Path for config casc file
ENV CASC_JENKINS_CONFIG /var/jenkins_home/casc.yaml

#Installing Jenkins Plugins
COPY plugins.txt /usr/share/jenkins/ref/plugins.txt
RUN /usr/local/bin/install-plugins.sh < /usr/share/jenkins/ref/plugins.txt

#Config file for jenkins - manages authentication ,  authorization , Jenkins URL , Creating users , etc
COPY casc.yaml /var/jenkins_home/casc.yaml


# Run this command to run jenkins
#  You can also mention the username and password in th casc.yaml file. 