CREATE TABLE IF NOT EXISTS medicines (
	id TEXT PRIMARY KEY,
	name TEXT NOT NULL UNIQUE,
	description TEXT,
	price REAL NOT NULL,
	stock INTEGER NOT NULL DEFAULT 0,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS visits (
	id TEXT PRIMARY KEY,
	patient_id TEXT NOT NULL,
	status TEXT NOT NULL DEFAULT 'registered',
	queue_number INTEGER NOT NULL,
	symptoms TEXT NOT NULL,
	diagnosis TEXT,
	prescription TEXT,
	registered_at DATETIME NOT NULL,
	examined_at DATETIME,
	dispensed_at DATETIME,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
	FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_visits_patient_id ON visits(patient_id);
CREATE INDEX IF NOT EXISTS idx_visits_status ON visits(status);

-- Insert sample medicines
INSERT OR IGNORE INTO medicines VALUES
('med-001', 'Paracetamol', 'Fever & pain reliever', 5000, 100, datetime('now'), datetime('now')),
('med-002', 'Amoxicillin', 'Antibiotic', 15000, 50, datetime('now'), datetime('now')),
('med-003', 'Ibuprofen', 'Anti-inflammatory', 8000, 75, datetime('now'), datetime('now'));

