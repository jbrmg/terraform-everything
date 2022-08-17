package de.jb.tfeverything;

import de.jb.tfeverything.repository.CabinetEntity;
import de.jb.tfeverything.repository.CounterTopEntity;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.rest.core.config.RepositoryRestConfiguration;
import org.springframework.data.rest.webmvc.config.RepositoryRestConfigurer;
import org.springframework.web.servlet.config.annotation.CorsRegistry;

@SpringBootApplication
public class KitchenApplication implements RepositoryRestConfigurer {

    public static void main(String[] args) {
        SpringApplication.run(KitchenApplication.class, args);
    }

    @Override
    public void configureRepositoryRestConfiguration(RepositoryRestConfiguration config, CorsRegistry cors) {
        config.exposeIdsFor(CabinetEntity.class, CounterTopEntity.class);
    }
}
