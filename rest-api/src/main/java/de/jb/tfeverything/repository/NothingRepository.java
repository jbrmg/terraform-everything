package de.jb.tfeverything.repository;

import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

import java.util.UUID;

@RepositoryRestResource(collectionResourceRel = "nothings", path = "nothings")
public interface NothingRepository extends PagingAndSortingRepository<NothingEntity, UUID> {
}
