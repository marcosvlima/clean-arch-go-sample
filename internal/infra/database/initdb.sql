-- Cria a tabela 'orders' para armazenar informações sobre pedidos.
CREATE TABLE IF NOT EXISTS orders (
    -- ID único do pedido (UUID ou string alfanumérica é comum).
    id VARCHAR(255) NOT NULL,
    -- Preço total dos itens no pedido, antes de impostos e taxas.
    price DECIMAL(10, 2) NOT NULL,
    -- Valor total de impostos (ex: ICMS, IPI, IVA) aplicados ao pedido.
    tax DECIMAL(10, 2) NOT NULL,
    -- Preço final que o cliente paga (price + tax).
    final_price DECIMAL(10, 2) NOT NULL,
    -- Data e hora em que o pedido foi criado (útil para rastreamento).
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- Chave primária para identificação única.
    PRIMARY KEY (id)
);

---

-- Insere exemplos de dados de pedidos (orders) mais realistas.
INSERT INTO
    orders (id, price, tax, final_price)
VALUES
    -- Exemplo 1: Pedido de E-commerce (Preço: 150.00, Imposto: 15.00)
    ('ord-abc-12345', 150.00, 15.00, 165.00),
    -- Exemplo 2: Pedido de Software/Serviço (Preço: 299.99, Imposto: 30.00)
    ('ord-def-67890', 299.99, 30.00, 329.99),
    -- Exemplo 3: Pedido com desconto aplicado (Preço base: 50.75, Imposto: 5.07)
    ('ord-ghi-11223', 50.75, 5.07, 55.82),
    -- Exemplo 4: Pedido de alto valor (Preço: 899.50, Imposto: 90.05)
    ('ord-jkl-44556', 899.50, 90.05, 989.55),
    -- Exemplo 5: Pedido de valor menor (Preço: 19.90, Imposto: 1.99)
    ('ord-mno-77889', 19.90, 1.99, 21.89)
-- A cláusula ON DUPLICATE KEY UPDATE é mantida para garantir a idempotência
-- (execução segura em caso de reexecução do script de inicialização).
ON DUPLICATE KEY UPDATE
    id = VALUES(id); -- Apenas atualiza o 'id' com seu próprio valor, essencialmente fazendo nada, mas evitando erro de duplicidade.