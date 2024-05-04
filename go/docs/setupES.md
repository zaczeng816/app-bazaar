```bash
sudo apt install default-jre 
sudo apt install apt-transport-https
# add the apt keys to so that apt can find the download link for es
wget -qO- https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add - 
sudo sh -c 'echo "deb  https://artifacts.elastic.co/packages/7.x/apt stable main" > /etc/apt/sources.list.d/elastic-7.x.list' 
sudo apt update
sudo apt install elasticsearch
# setup the configuration of elasticsearch
sudo vim /etc/elasticsearch/elasticsearch.yml
sudo systemctl enable elasticsearch
sudo systemctl start  elasticsearch
# add a test user to es
sudo /usr/share/elasticsearch/bin/elasticsearch-users useradd $NEW_USER_NAME -p $NEW_PASSWORD -r superuse
```