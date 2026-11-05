package com.tuca.order_service.service;

import java.util.Arrays;
import java.util.List;
import java.util.UUID;

import org.springframework.stereotype.Service;

import com.tuca.order_service.dto.InventoryResponse;
import com.tuca.order_service.dto.OrderLineItemsDto;
import com.tuca.order_service.dto.OrderRequest;
import com.tuca.order_service.model.Order;
import com.tuca.order_service.model.OrderLineItems;
import com.tuca.order_service.repository.OrderRepository;
import org.springframework.web.reactive.function.client.WebClient;

import jakarta.transaction.Transactional;
import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
@Transactional
public class OrderService {
    
    private final OrderRepository orderRepository;
    private final WebClient webClient;

    public void placeOrder(OrderRequest orderRequest) {
        Order order = new Order();
        order.setOrderNumber(UUID.randomUUID().toString());

        List<OrderLineItems> orderLineItemsList = orderRequest.getOrderLineItemsDtoList()
            .stream()
            .map(this::mapToDto)
            .toList();

        order.setOrderLineItems(orderLineItemsList);
 
        // Call Inventory Service and place order if product is in stock
        List<String> skuCodes = order.getOrderLineItems().stream()
            .map(OrderLineItems::getSkuCode)
            .toList();

        InventoryResponse[] inventoryResponses = webClient.get()
            .uri("http://localhost:8082/api/v1/inventory?skuCode=", uriBuilder -> uriBuilder.queryParam("skuCode", skuCodes).build())
            .retrieve()
            .bodyToMono(InventoryResponse[].class)
            .block();

        boolean allProductsInStock = Arrays.stream(inventoryResponses).allMatch(InventoryResponse::isInStock);

        if (allProductsInStock) {
            orderRepository.save(order);
        }
        else {
            throw new IllegalArgumentException("A product is not in stock, please try again later.");
        }
    }

    private OrderLineItems mapToDto(OrderLineItemsDto orderLineDto) {
        OrderLineItems orderLineItems = new OrderLineItems();
        orderLineItems.setPrice(orderLineDto.getPrice());
        orderLineItems.setQuantity(orderLineDto.getQuantity());
        orderLineItems.setSkuCode(orderLineDto.getSkuCode());

        return orderLineItems;
    }
}
