IF OBJECT_ID('dbo.Ticket', 'U') IS NOT NULL
DROP TABLE dbo.Ticket
GO
CREATE TABLE dbo.Ticket
(
    Id INT NOT NULL PRIMARY KEY,
);
GO


IF OBJECT_ID('dbo.WinningTickets', 'U') IS NOT NULL
DROP TABLE dbo.WinningTickets
GO
CREATE TABLE dbo.WinningTickets
(
    WinningTicketsId INT NOT NULL PRIMARY KEY,
    TicketId [INT] FOREIGN KEY REFERENCES Ticket(Id)
);
GO

IF OBJECT_ID('dbo.TicketNumbers', 'U') IS NOT NULL
DROP TABLE dbo.TicketNumbers
GO
CREATE TABLE dbo.TicketNumbers
(
    TicketNumbersId INT NOT NULL PRIMARY KEY,
    TicketId [INT] FOREIGN KEY REFERENCES Ticket(Id)
);
GO

BULK INSERT Ticket
from 'C:\Source\repos\advent-of-code-2023\day-4-sql\tickets.csv'
with (firstrow = 1,
      fieldterminator = ',',
      rowterminator='\n',
      batchsize=10000,
      maxerrors=10);


SELECT *
FROM Ticket;