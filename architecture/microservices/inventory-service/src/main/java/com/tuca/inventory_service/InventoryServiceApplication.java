package com.tuca.inventory_service;

import com.tuca.inventory_service.repository.InventoryRepository;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.CommandLineRunner;
import org.springframework.context.annotation.Bean;
import com.tuca.inventory_service.model.Inventory;
import lombok.RequiredArgsConstructor;


@SpringBootApplication
@RequiredArgsConstructor
public class InventoryServiceApplication {

    private final InventoryRepository inventoryRepository;

	public static void main(String[] args) {
		SpringApplication.run(InventoryServiceApplication.class, args);
	}

	@Bean
	public CommandLineRunner loadData(InventoryRepository inventoryRepository) {
		return args -> {
			Inventory inventory = new Inventory();
			inventory.setSkuCode("switch_2_base");
			inventory.setQuantity(100);
			
			Inventory inventory2 = new Inventory();
			inventory2.setSkuCode("switch_2_mario_kart");
			inventory2.setQuantity(50);

			inventoryRepository.save(inventory);
			inventoryRepository.save(inventory2);
		};
	}
}
