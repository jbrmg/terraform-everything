package de.jb.tfeverything.repository;

import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

import java.util.UUID;

@RepositoryRestResource(collectionResourceRel = "cabinets", path = "cabinets")
public interface CabinetRepository extends PagingAndSortingRepository<CabinetEntity, UUID> {
}
