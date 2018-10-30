#!/bin/bash env
# description: install monit and auto config for blemobi micro service
# auth: yanue
# time: 2018-10-30
# usage:
# > wget https://raw.githubusercontent.com/yanue/go-esport-common/master/bin/monit.sh
# > chmod +x monit.sh
# > ./monit.sh

export PATH=/usr/local/bin:/sbin:/usr/bin:/bin

disableYumPlugin(){
    echo -e "\033[32m1. set yum plugin disable \033[0m"
    sed -i -e 's/plugins.*/plugins=0/' /etc/yum.conf
    echo $0

    sed -i 's/.*\(alias vi\).*/\1/' /etc/bashrc
}

installBase(){
    echo -e "\033[32m2. install base \033[0m"
    yum install -y vim
}

setBashrc(){
    echo -e "\033[32m3. set bashrc \033[0m"
    # alias vi="vim"
    if !(grep -R "alias\s*vi" /etc/bashrc)
    then
        echo 'alias vi="vim"' >> /etc/bashrc
    fi

    # alias nets="netstat -ntpl"
    if !(grep -R "alias\s*nets" /etc/bashrc)
    then
        echo 'alias nets="netstat -ntpl"' >> /etc/bashrc
    fi

    # alias j="cd /opt/miwu/service/"
    if !(grep -R "alias\s*j" /etc/bashrc)
    then
        echo 'alias j="cd /opt/miwu/service/"' >> /etc/bashrc
    fi

    # source
    source /etc/bashrc
}

monitInstall(){
    echo -e "\033[32m4. install monit \033[0m"
    yum install -y monit
}

monitEnv(){
    echo -e "\033[32m5. change monit config \033[0m"
    text=$(cat <<-END
###############################################################################
## Global section
###############################################################################

set daemon  10              # check services at 30 seconds intervals
set log syslog
set httpd port 2812 and
    use address localhost  # only accept connection from localhost
    allow localhost        # allow localhost to connect to the server and
    allow admin:monit

###############################################################################
## Includes
###############################################################################

# set daemon mode timeout to 10 seconds
set daemon 10
# Include all files from /etc/monit.d/
include /etc/monit.d/*
END
)
    echo "$text" > /etc/monitrc
}

monitConfConsul(){
    echo -e "\033[32m6. generating consul conf \033[0m"
text=$(cat <<-END
check process consul
    matching "consul"
    start program = "/bin/systemctl start consul"
    stop program = "/bin/systemctl stop consul"
    if failed port 8500 use type tcp then restart
END
)
    echo "$text" > /etc/monit.d/consul.conf
}

monitConfMicro(){
    echo -e "\033[32m7. generating blemobi micro service conf \033[0m"

    arr=($(netstat -ntpl | grep 'blemobi' | awk '{print $4"|"$7}'))

    declare -A nameWithPort=()

    for txt in "${arr[@]}"
    do
        name=$(echo $txt | awk '{split($0,a,"|" ); print a[2]}' | awk '{split($0,b,"/" ); print b[2]}')
        port=$(echo $txt | awk '{split($0,a,"|" ); print a[1]}' | awk '{n=split($1,A,":"); print A[n]}')

        if [[ $name == "" || $port == "" ]]; then
            echo "process name or port not found:",$name,$port
            continue
        fi

        # replace min port on same name
        needReplace=0
        if [[ ${nameWithPort[$name]} != "" ]];
        then
            # if exist port ge now port
            if test ${nameWithPort["$name"]} -ge $port; then
                needReplace=1
            fi
        else
            # if not exist
            needReplace=1
        fi

        if [ "$needReplace" = "1" ]; then
            nameWithPort["$name"]=$port
        fi
    done

    for runName in ${!nameWithPort[@]}
    do
        # some var
        serviceName=$(echo $runName | sed -e "s/-//")
        monitName=$(echo $runName | sed -e "s/-//"| sed -e "s/blemobi//")
        runPort=${nameWithPort[$runName]}

        # check systemd file
        if [ -f "/etc/systemd/system/$serviceName.service" ]
        then
            # monit data
            text=$(cat <<-END
check process $monitName
    matching "$runName"
    start program = "/bin/systemctl start $serviceName"
    stop program = "/bin/systemctl stop $serviceName"
    if failed port $runPort use type tcp then restart
END
)
            # 30-37m 黑红绿黄蓝紫天白
            echo -e "\033[34m> monit conf:\033[0m"
            echo -e "\033[33m$text\033[0m"
            echo -e "\033[34m> /etc/monit.d/$monitName.conf\033[0m"
            echo "$text" > "/etc/monit.d/$monitName.conf"
        fi
    done
}

monitReload(){
    echo -e "\033[32m8. start monit \033[0m"
    /bin/systemctl restart monit
    /usr/bin/monit reload
    /usr/bin/monit status
}

init(){
    disableYumPlugin
    installBase
    setBashrc
    monitInstall
    monitEnv
    monitConfConsul
    monitConfMicro
    monitReload
    echo -e "\033[32mdone\033[0m"
}

init