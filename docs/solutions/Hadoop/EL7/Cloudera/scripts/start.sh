#!/bin/bash
## First Script called on Bastion Host after TF Deployment by remote-exec

echo -e "Installing Screen on Bastion"
sudo yum install screen.x86_64 -y
sudo sed -i 's/1000/10000/g' /etc/screenrc
sudo tee -a ~/.screenrc << EOF
screen -t setup 
select 0
screen -t logwatch
select 1

altscreen on
term screen-256color
bind ',' prev
bind '.' next

hardstatus alwayslastline
hardstatus string '%{= kG}[ %{G}%H %{g}][%= %{= kw}%?%-Lw%?%{r}(%{W}%n*%f%t%?(%u)%?%{r})%{w}%?%+Lw%?%?%= %{g}][%{B} %m-%d %{W}%c %{g}]'
EOF
sudo tee -a /etc/screenrc << EOF
caption always "%{= bb}%{+b w}%n %t %h %=%l %H %c"
hardstatus alwayslastline "%-Lw%{= BW}%50>%n%f* %t%{-}%+Lw%<"
activity "Activity in %t(%n)"

shelltitle "shell"
shell -$SHELL
EOF
echo -e "Starting Master Cluster Provisioning Process"
sudo screen -dmLS bastion 
sleep .001
## Start Bastion setup script, time it so the cluster build time is tracked
sudo screen -S bastion -t setup -X stuff '/usr/bin/time /home/opc/bastion.sh\n'
