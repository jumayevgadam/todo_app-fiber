-- For multiple subcategories, we must add parent_id
ALTER TABLE Categories
    ADD COLUMN parent_id BIGINT;