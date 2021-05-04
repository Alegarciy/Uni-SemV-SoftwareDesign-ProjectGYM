USE PlusGymProject ;
GO  
CREATE VIEW dbo.CompleteSessions AS  
    SELECT 
        [session].Id                            AS SessionID,
        [session].Fecha                         AS SessionDate,
        [session].Cancelada                     AS IsCancelled,
        
        [preliminarySession].Costo              AS Cost,
        [preliminarySession].Cupo               AS Spaces,
        [preliminarySession].DuracionMinutos    AS Duration,
        [preliminarySession].HoraInicio         AS StartTime,
        
        [service].Id                            AS ServiceId,
        [service].Nombre                        AS ServiceName,

        [instructor].Id                         AS InstructorId,
        [instructor].Nombre                     AS InstructorName,
        [instructor].Cedula                     AS InstructorIdentification,
        [instructor].Correo                     AS InstructorEmail,

        [instructorType].Id                     AS InstructorTypeId,
        [instructorType].Nombre                 AS InstructorType,

        [room].Id                               AS RoomId,
        [room].Nombre                           AS RoomName
    FROM 
        dbo.Sesion AS [session]
    INNER JOIN 
        dbo.SesionPreliminar AS preliminarySession
        ON [preliminarySession].Id = [session].SessionPreliminarId
    INNER JOIN 
        dbo.Especialidades AS [service]
        ON [service].Id = [preliminarySession].EspecialidadId
    INNER JOIN
        dbo.Instructor AS instructor
        ON instructor.Id = preliminarySession.InstructorId
    INNER JOIN
        dbo.TipoInstructor AS instructorType
        ON instructor.Tipo = instructorType.Id
    INNER JOIN 
        dbo.Sala AS room 
        ON room.Id = preliminarySession.SalaId
    ;
GO
-- Query to test
SELECT *  
FROM dbo.CompleteSessions
