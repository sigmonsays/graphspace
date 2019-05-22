graphspace - graphviz workspace tool for making dot graphs

- simple standalone web server 
- validates graphviz syntax
- saves graphs in sqlite for easy sharing
- configurable output image size
- all graph layouts supported
- simple on disk cache to improve load time

[Screenshot is worth a thousand words](https://raw.githubusercontent.com/sigmonsays/graphspace/master/static/graphspace.jpg)

# docker

docker hub image is built at https://hub.docker.com/r/sigmonsays/graphspace


launch a graphspace container

    docker run -d --name graphspace1 sigmonsays/graphspace 

connect directly to private docker ip

    IP="$(docker inspect graphspace1 -f '{{ .NetworkSettings.IPAddress }}')"
    echo url is http://$IP:7001


launch a graphspace container and expose port

    docker run -d -p 7001:7001 --name graphspace1 sigmonsays/graphspace 


