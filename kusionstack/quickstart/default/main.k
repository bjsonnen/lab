import kam.v1.app_configuration as ac
import service
import service.container as c
import network as n
import mysql

# main.k declares the customized configuration codes for default stack. 
quickstart: ac.AppConfiguration {
    workload: service.Service {
        containers: {
            quickstart: c.Container {
                image: "kusionstack/kusion-quickstart:latest"
                env: {
                    "DB_HOST": "$(KUSION_DB_HOST_QUICKSTART_DEFAULT_QUICKSTART_MYSQL)"
                    "DB_USERNAME": "$(KUSION_DB_USERNAME_QUICKSTART_DEFAULT_QUICKSTART_MYSQL)"
                    "DB_PASSWORD": "$(KUSION_DB_PASSWORD_QUICKSTART_DEFAULT_QUICKSTART_MYSQL)"
                }
            }
        }
    }
    accessories: {
        "network": n.Network {
            ports: [
                n.Port {
                    port: 8080
                }
            ]
        }
        "mysql": mysql.MySQL {
            type: "local"
            version: "8.0"
        }
    }
}
