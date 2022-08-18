package de.jb.tfeverything.repository;

import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

import java.util.Optional;
import java.util.UUID;

@RepositoryRestResource(collectionResourceRel = "countertops", path = "countertops")
public interface CounterTopRepository extends PagingAndSortingRepository<CounterTopEntity, UUID> {

    Optional<CounterTopEntity> findByCabinetIdsContains(UUID cabinetId);

}
