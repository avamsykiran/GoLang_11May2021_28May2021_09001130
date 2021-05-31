RDBMS
------------------------------------------------

    Relational Database Management System:

        Relation
                    is a mapping from one set to another
                    and these sets are called
                        Domain and Range.


        Relation: f:Domain--->Range

        Emps
        --------------------------------------------------
        EmpId   Ename   Sal     Desgn
        --------------------------------------------------
        101     Vamsy   4567    Sr.Analyst


        Database
                Table / Relations
                    Records / Rows / Tuples
                    Attributes / Cols  / Fields


        Normalization and Denormalization
            to construct a design a RDBMS database.

        RDBMS Products
        ----------------------------------------------------
            Oracle
            Ms SQL Server
        
            MySQL
            PostgreSQL


        PostgresSQL
        --------------------------------------------------------

            Postgre Server
            psql    Command Line Interaction Tool
            pgAdmin GUI Tool for interaction


            psql
            -----------------------------------------------------

                \l          List all Databases
                \c          Connect to a database
                \d          List all database obejct like tables,view...etc
                \q          To quit from psql


            SQL on Postgre
            ----------------------------------------------------------

                DDL     Data Definition Language

                            CREATE      ALTER       DROP        TRUNCATE

                DML     Data Manipulation Language

                            INSERT      UPDATE      DELETE

                DRL     Data Retrival Language

                            SELECT
                                FROM
                                WHERE
                                GROUP BY
                                HAVING
                                ORDER By

                            Aggregation Function
                            Sub Quereis
                            Co-Related Sub Queries
                            Joins

                DCL     Data Control Language

                            GRANT               REVOKE

                TCL     Transaction Control Language

                            SAVEPOINT           COMMIT          ROLLBAck

        Data Definition Language
        ---------------------------------------------

        CREATE DATABASE database-name;
        DROP DATABASE database-name;

        CREATE TABLE tableName(
            colName datatype constraint,
            colName datatype constraint,
            colName datatype constraint,
            colName datatype constraint,
            .....................
            colName datatype constraint
        );

        DROP TABLE tableName;
        TRUNCATE tableName;

        Constraint
                is a rule assaigned to assure data integrity (validity)

                Domain Level Constraint         Domain Integrity

                    Data Types  
                        Numeric
                            int
                            smallint
                            auto_increment
                            float
                            numric
                            float8

                        Non Numeric
                            char(n)
                            varchar(n)
                            text

                            date
                            time
                            timestamp

                            array
                            json

                    not null                    
                    check
                    default

                Entity Level Constraint         Entity Integrity

                    unique
                    check
                    primary key
                    foreign key


            create table contact_groups(
                grp_id  int primary key,
                grp_name text not null
            );

            create table contacts(
                cid            int      primary key,
                contact_name   text     not null,
                mobile         char(10) unique,
                dob            date     not null,
                mailId         text     unique,
                grp_id         int,
                constraint fk_contact_groups_contacts foreign key(grp_id) references contact_groups(grp_id)
            );


            Vehicle Servecing center,
                Keep track of the vehicles and their servicing records:
                    vehicleRegNumber, modela dn make of the vehicle, category of the vehicle (SUV/MUV...),
                    serviceId, service number, complaint, solution, charge/Fee, dateOfService and the vehicle owner
                    details like ownerId, name, contactNumber.

                    A owner can have many vehicles
                    a vehicle can be owned by different owners
                    a vehcile will be serviced many times,
                    each service addresses many complaints.

                                        PK
                Owners              owner_id           name,mobile
                Vehicles            reg_num            model,make,category
                Vehcile_Owners      owner_Id,reg_num   reg_date
                Services            service_id         service_date,reg_num
                Complaints          complaintId        complaint,solution,fee,service_id


        ALTER TABLE tableName   ADD colName datatype constraint
        ALTER TABLE tableName   DROP colName
        ALTER TABLE tableName   MODIFY colName datatype constraint
        ALTER TABLE tableName   ALTER COLUMN colName TYPE datatype
        ALTER TABLE tableName   ADD constraint constraint_name constraintDefinition
        ALTER TABLE tableName   DROP constraint constraint_name


        alter table contact_groups add colxyz numeric;
        alter table contact_groups alter column colxyz type int;
        alter table contact_groups modify colxyz int not null;
        alter table contact_groups drop column colxyz;

        alter table contact_groups add constraint unq_contacts_groups UNIQUE(grp_name);
        alter table contact_groups drop constraint unq_contacts_groups;

        


           

