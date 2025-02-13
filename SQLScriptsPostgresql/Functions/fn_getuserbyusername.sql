-- =============================================
-- Author:      Jorge Gutiérrez Cordero (PostgreSQL adaptation)
-- Create Date: Current
-- Description: Gets a user by its username
-- =============================================

CREATE OR REPLACE FUNCTION FN_GetUserByUsername(
    p_username VARCHAR(20)
) RETURNS TABLE (
    unique_identifier INT,
    username VARCHAR(20),
    password VARCHAR(255),
    user_type INT
) AS $$
BEGIN
    -- Return user information
    RETURN QUERY
    SELECT 
        usr.unique_identifier,
        usr.username,
        usr.password,
        usr.type_id AS user_type
    FROM 
        complete_users usr
    WHERE 
        usr.username = p_username
        AND usr.active = true;
END;
$$ LANGUAGE plpgsql;

-- Example usage:
-- SELECT * FROM FN_GetUserByUsername('Cliente1'); 