INSERT INTO users (id, name, identification_number, email, phone, created_at) VALUES
  ('79f1ed5d-fb02-480f-8430-caabf42ce630', 'Juan Pérez', '123', 'juan.perez@example.com', '123456789', CURRENT_TIMESTAMP),
  ('1155b711-503a-48c5-922c-8c8ed4c3e3c3', 'María Gómez', '456','maria.gomez@example.com', '987654321', CURRENT_TIMESTAMP),
  ('08644620-133a-411a-988c-3f714fec5020', 'Carlos López', '789','carlos.lopez@example.com', '1122334455', CURRENT_TIMESTAMP);
--
INSERT INTO merchants (id, name, description, created_at) VALUES
  ('ee15da8e-6386-40e2-ba89-24bff3234e4b', 'Texaco', 'Estación de servicio de gasolina', CURRENT_TIMESTAMP),
  ('e47a9c8e-9004-4566-90ec-23a8ae149f63', 'Supermercados Leal', 'Cadena de supermercados', CURRENT_TIMESTAMP);
--
INSERT INTO branches (id, merchant_id, name, location, created_at) VALUES
  ('51968d08-9018-4a5c-94e6-9c7181359ef7', 'ee15da8e-6386-40e2-ba89-24bff3234e4b', 'Texaco Sucursal 1', 'Calle 10 #45-67, Bogotá', CURRENT_TIMESTAMP),
  ('39afa126-5a23-486c-ad36-1825a111e348', 'ee15da8e-6386-40e2-ba89-24bff3234e4b', 'Texaco Sucursal 2', 'Carrera 15 #23-56, Medellín', CURRENT_TIMESTAMP),
  ('4f747ef6-0a13-4af8-8bab-65f9cf580836', 'e47a9c8e-9004-4566-90ec-23a8ae149f63', 'Supermercado Leal Centro', 'Avenida Central #12-34, Bogotá', CURRENT_TIMESTAMP);
--
INSERT INTO campaigns (id, branch_id, start_date, end_date, bonus_type, bonus_value, min_purchase, created_at) VALUES
  ('bdbc9716-217b-427e-80d0-a1e5f09bd3c4', '51968d08-9018-4a5c-94e6-9c7181359ef7', '2024-05-15', '2024-05-30', 'double', 2.0, 0, CURRENT_TIMESTAMP), -- Doble de puntos o cashback sin monto mínimo
  ('1434f164-627a-4036-920e-2be28dc00bee', '39afa126-5a23-486c-ad36-1825a111e348', '2024-05-15', '2024-05-20', 'percentage', 0.30, 20000, CURRENT_TIMESTAMP); -- 30% adicional para compras mayores a $20,000
--
INSERT INTO rewards (id, user_id, branch_id, reward_type, reward_value, created_at) VALUES
  ('a5ca5b06-7fbb-42e8-8100-356e99e1eb45', '79f1ed5d-fb02-480f-8430-caabf42ce630', '51968d08-9018-4a5c-94e6-9c7181359ef7', 'points', 100, CURRENT_TIMESTAMP),  -- Juan Pérez tiene 100 puntos en Texaco Sucursal 1
  ('4b060ed1-e462-4879-8fd5-1f3687030ccc', '1155b711-503a-48c5-922c-8c8ed4c3e3c3', '39afa126-5a23-486c-ad36-1825a111e348', 'cashback', 50, CURRENT_TIMESTAMP), -- María Gómez tiene 50 en cashback en Texaco Sucursal 2
  ('932a0f09-60fd-480c-b89c-7e02eebef348', '08644620-133a-411a-988c-3f714fec5020', '4f747ef6-0a13-4af8-8bab-65f9cf580836', 'points', 200, CURRENT_TIMESTAMP);  -- Carlos López tiene 200 puntos en Supermercado Leal
--
INSERT INTO purchases (user_id, branch_id, purchase_amount, reward_earned, reward_type, campaign_id, created_at) VALUES
  ('79f1ed5d-fb02-480f-8430-caabf42ce630', '51968d08-9018-4a5c-94e6-9c7181359ef7', 20000, 40, 'points', 'bdbc9716-217b-427e-80d0-a1e5f09bd3c4', CURRENT_TIMESTAMP), -- Juan Pérez compró por 20,000 y ganó 40 puntos (por campaña)
  ('1155b711-503a-48c5-922c-8c8ed4c3e3c3', '39afa126-5a23-486c-ad36-1825a111e348', 25000, 32.5, 'cashback', '1434f164-627a-4036-920e-2be28dc00bee', CURRENT_TIMESTAMP), -- María Gómez compró por 25,000 y ganó 32.5 cashback (por campaña)
  ('08644620-133a-411a-988c-3f714fec5020', '4f747ef6-0a13-4af8-8bab-65f9cf580836', 30000, 60, 'points', NULL, CURRENT_TIMESTAMP); -- Carlos López compró por 30,000 y ganó 60 puntos (sin campaña)
