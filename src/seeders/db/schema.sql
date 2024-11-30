-- Create `users` table
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

-- Create `orders` table without foreign key
CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add foreign key constraint for `orders`
ALTER TABLE orders
ADD CONSTRAINT fk_user FOREIGN KEY (user_id)
REFERENCES users(id) ON DELETE CASCADE;
