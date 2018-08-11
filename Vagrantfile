Vagrant.configure("2") do |config|
    config.vm.box = "ubuntu/xenial64"
    config.vm.provider "virtualbox" do |vb|
        vb.cpus   = 1
        vb.memory = 1024
      end
    config.vm.provision "ansible" do |ansible|
        ansible.playbook = "ansible/ansible-docker.yml"
        ansible.extra_vars = { 
            ansible_python_interpreter:"/usr/bin/python3",
            OPENWEATHER_API_KEY: ENV['OPENWEATHER_API_KEY'],
            CITY_NAME: ENV['CITY_NAME']
         }
    end
  end