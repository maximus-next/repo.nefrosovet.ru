-- +goose Down
DROP TABLE IF EXISTS appointment CASCADE;
DROP TABLE IF EXISTS appointment_param CASCADE;
DROP TABLE IF EXISTS appointment_program CASCADE;
DROP TABLE IF EXISTS clinic CASCADE;
DROP TABLE IF EXISTS confirm_code CASCADE;
DROP TABLE IF EXISTS employee CASCADE;
DROP TABLE IF EXISTS employee CASCADE;
DROP TABLE IF EXISTS implementation CASCADE;
DROP TABLE IF EXISTS implementation_param CASCADE;
DROP TABLE IF EXISTS invite CASCADE;
DROP TABLE IF EXISTS patient CASCADE;