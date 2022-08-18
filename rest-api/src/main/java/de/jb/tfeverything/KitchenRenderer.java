package de.jb.tfeverything;

import de.jb.tfeverything.repository.CabinetEntity;
import de.jb.tfeverything.repository.CabinetRepository;
import de.jb.tfeverything.repository.CounterTopRepository;
import de.jb.tfeverything.repository.KitchenEntity;
import de.jb.tfeverything.repository.KitchenRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.ServletContext;
import java.io.IOException;
import java.io.InputStream;
import java.util.Optional;
import java.util.Set;
import java.util.UUID;

@RestController
@RequiredArgsConstructor
public class KitchenRenderer {

    private final ServletContext servletContext;
    private final KitchenRepository kitchenRepository;
    private final CabinetRepository cabinetRepository;
    private final CounterTopRepository counterTopRepository;

    @GetMapping(path = "/render/{id}")
    public ResponseEntity renderKitchen(@PathVariable("id") UUID kitchenId) throws IOException {
        Optional<KitchenEntity> kitchen = kitchenRepository.findById(kitchenId);
        if (!kitchen.isPresent()) {
            return ResponseEntity.notFound().build();
        }

        Set<CabinetEntity> cabinets = cabinetRepository.findByKitchenId(kitchenId);
        if (cabinets.isEmpty()) {
            return writeImageResponse("kitchen.jpg");
        }

        boolean counterTopPresent = cabinets.stream()
                .map(c -> counterTopRepository.findByCabinetIdsContains(c.getId()))
                .allMatch(ct -> ct.isPresent());

        if (!counterTopPresent) {
            return writeImageResponse("cabinets.jpg");
        }

        return writeImageResponse("countertop.jpg");

    }

    private ResponseEntity<byte[]> writeImageResponse(String imageName) throws IOException {
        InputStream in = getClass().getResourceAsStream("/static/" + imageName);
        return ResponseEntity.ok().contentType(MediaType.IMAGE_JPEG)
                .body(in.readAllBytes());
    }
}
