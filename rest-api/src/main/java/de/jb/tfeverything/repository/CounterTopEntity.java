package de.jb.tfeverything.repository;

import lombok.Data;

import javax.persistence.ElementCollection;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import java.util.List;
import java.util.UUID;

@Entity
@Data
public class CounterTopEntity {

    public enum CounterTopType {
        EKBACKEN,
        KARLBY,
        SKARARP
    }

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private UUID id;
    private CounterTopType type;
    @ElementCollection
    private List<UUID> cabinetIds;

}
