package de.jb.tfeverything.repository;

import lombok.Data;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import java.util.UUID;

@Entity
@Data
public class CabinetEntity {

    public enum FrontType {
        RINGHULT,
        VOXTORP,
        VEDDINGE
    }

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private UUID id;
    private String color;
    private FrontType front;

}
