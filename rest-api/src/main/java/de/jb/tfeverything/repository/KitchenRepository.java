package de.jb.tfeverything.repository;

import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

import java.util.UUID;

@RepositoryRestResource(collectionResourceRel = "kitchens", path = "kitchens")
public interface KitchenRepository extends PagingAndSortingRepository<KitchenEntity, UUID> {
}
