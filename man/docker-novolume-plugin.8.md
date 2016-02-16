% DOCKER-NOVOLUME-PLUGIN(8)
% Antonio Murdaca
% FEBRUARY 2016
# NAME
docker-novolume-plugin - Blocks self provisioned volumes

# SYNOPSIS
**docker-novolume-plugin**
[**--host**=[=*unix:///var/run/docker.sock*]]

# DESCRIPTION
When a volume in provisioned via the VOLUME instruction in a Dockerfile or via
docker run -v volumename, host's storage space is used. This could lead to an
unexpected out of space issue which could bring down everything. There are situations
where this is not an accepted behavior. PAAS, for instance, can't allow their users
to run their own images without the risk of filling the entire storage space on a server.
One solution to this is to deny users from running images with volumes. This way the
only storage a user gets can be limited and PAAS can assign quota to it.
This plugin solves this issue by disallowing starting a container with local volumes defined. In particular, the plugin will block docker run with:

    --volumes-from
    images that have VOLUME(s) defined
    volumes early provisioned with docker volume command

The only thing allowed will be just bind mounts.

# OPTIONS

**--host**="unix:///var/run/docker.sock"
  Specifies the host where to contact the docker daemon.

# AUTHORS
Antonio Murdaca <runcom@redhat.com>
