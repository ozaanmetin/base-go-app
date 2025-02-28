CREATE TABLE regions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    name VARCHAR(255) NOT NULL,
    country_id UUID NOT NULL,
    parent_id UUID,

    CONSTRAINT fk_country_id FOREIGN KEY (country_id) REFERENCES countries(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_parent_id FOREIGN KEY (parent_id) REFERENCES regions(id) ON DELETE SET NULL ON UPDATE CASCADE
);