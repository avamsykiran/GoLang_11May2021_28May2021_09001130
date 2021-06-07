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
                            Sub Queries
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
        ALTER TABLE tableName   ALTER COLUMN colName TYPE datatype
        ALTER TABLE tableName   ADD constraint constraint_name constraintDefinition
        ALTER TABLE tableName   DROP constraint constraint_name


        alter table contact_groups add colxyz numeric;
        alter table contact_groups alter column colxyz type int;
        alter table contact_groups drop column colxyz;

        alter table contact_groups add constraint unq_contacts_groups UNIQUE(grp_name);
        alter table contact_groups drop constraint unq_contacts_groups;

        Data Manipulation Language
        ---------------------------------------------

        INSERT INTO table_name
        VALUES(VAL1,VAL2,VAL3.....) //order of col in table definition

        INSERT INTO table_name(col1,col2,col3)
        VALUES(VAL1,VAL2,VAL3.....) //order of col in caommand

        UPDATE  table_name
        SET     col1=value,col2=value...
        [WHERE  cond];

        DELETE FROM table_name
        [WHERE  cond];
    
        Data Retrival Language
        ---------------------------------------------

            SELECT      *|COL1,COL2,COL3                //PROJECTION
            [ FROM       TABLE_AND_JOINS         ]        //DATA SOURCE
            [ WHERE      COND                    ]        //PRE-GROUPING FILTRATION
            [ GROUP BY   COL1,COL2,...           ]        //GROUPING
            [ HAVING     COND                    ]        //POST-GROUPING FILTRATION
            [ ORDER BY   COL1,COL2...            ]        //SORTING

            WHERE Clause - PRE-GROUPING FILTRATION
            ---------------------------------------

                =
                <
                >
                <>
                <=
                >=

                like
                is null
                is not null
                in 
                between .. and ..
                
                and
                or
                not

            Aggregate Functions
            -----------------------------------

                MIN         
                MAX
                SUM
                AVG
                COUNT

            Sub-Queries
            -------------------------------------

                select ename,hiredate
                from emp
                where hiredate = (select min(hiredate) from emp);

                select ename,hiredate,deptno
                from emp
                where hiredate = (select min(hiredate) from emp where deptno =10) and deptno = 10;


            Corelated Sub Queries
            -------------------------------------

                select e1.ename,e1.hiredate,e1.deptno
                from emp e1
                where e1.hiredate = (select min(e.hiredate) from emp e where e.deptno=e1.deptno);
           
            Group By Clause - GROUPING
            ---------------------------------------

                select min(hiredate),max(hiredate),sum(sal),min(sal),max(sal),deptno
                from emp
                group by deptno;

                select deptno,job,count(*)
                from emp
                group by deptno,job;

                select deptno,job,count(*)
                from emp
                group by job,deptno;

            Having Clause - Post-GRouping Filtration
            -------------------------------------------

                select deptno,job,count(*)
                from emp
                group by job,deptno
                having count(*)>=2;

            Order By Clause - Sorting
            --------------------------------------------

                select * from emp order by ename;
                select * from emp order by ename desc;
                select * from emp order by deptno,empno;

        Joins
        ====================================================================================

                is machanisim using which
                we can retrive related records
                from two or more tables!

                NATURAL JOIN

                        SELECT *|colNames
                        FROM table1 NATURAL JOIN table;

                CROSS JOIN

                        SELECT *|colNames
                        FROM table1 CROSS JOIN table;

                INNER JOIN

                        SELECT *|colNames
                        FROM table1 INNER JOIN table2 ON table1.col=table2.col;

                LEFT OUTER JOIN
        
                        SELECT *|colNames
                        FROM table1 LEFT OUTER JOIN table2 ON table1.col=table2.col;

                RIGHT OUTER JOIN

                        SELECT *|colNames
                        FROM table1 RIGHT OUTER JOIN table2 ON table1.col=table2.col;

                FULL OUTER JOIN

                        SELECT *|colNames
                        FROM table1 FULL OUTER JOIN table2 ON table1.col=table2.col;

        Postgree Built-in functions
        =====================================================================================
        
            string functions:
            ==================
            1) length(string)                          SELECT length('Vamsy Kiran');
            2) lower(string)
            3) upper(string)
            4) ltrim(string)
            5) rtrim(string)
            6) trim(string)
            7) lpad(string, length, pad_string)        SELECT lpad('Vamsy', 8, 'A');
            8) rpad(string, length, pad_string)
            9) initcap(string)
            10) substring( string [from start_position] [for length] )
                SELECT substring('Vamsy Kiran' from 1 for 4);
            11) repeat(string, number)
            12) replace( string, from_substring, to_substring )

            numeric/math functions:
            =======================
            1) abs(number)
            2) mod(n, m)
            3) power(m, n)
            6) sign(n)
            7) sqrt(n)
            8) ceiling(n)
            9) floor(n)
            10) round( number, [ decimal_places ] )
            11) trunc( number, [ decimal_places ] )

            date functions:
            ===============
            18) current_date
                SELECT current_date;

                SELECT current_date + 1;
            19) current_time
                SELECT current_time;
            20) now
                SELECT now();
            21) age( [date1,] date2 )

                SELECT age(timestamp '2014-01-01'); 	/* from current date */
                SELECT age(timestamp '2014-04-25', timestamp '2014-01-01');
                SELECT age(timestamp '2014-04-25', timestamp '2014-04-17');
                SELECT age(current_date, timestamp '2012-09-16');

            conversion functions:
            =====================
            22) to_char( value, format_mask )

                With numbers:
                -------------
                Parameter	Explanation
                ----------------------------
                9	Value (with no leading zeros)
                0	Value (with leading zeros)
                .	Decimal
                ,	Group separator
                PR	Negative value in angle brackets
                S	Sign
                L	Currency symbol
                D	Decimal
                G	Group separator
                MI	Minus sign (for negative numbers)
                PL	Plus sign (for positive numbers)
                SG	Plus/minus sign (for positive and negative numbers)
                RN	Roman numerals
                TH	Ordinal number suffix
                th	Ordinal number suffix
                V	Shift digits
                EEEE	Scientific notation

                    SELECT to_char(210, '0999.99');
                    0210.00

                    SELECT to_char(1210.7, '9G999.99');
                    1,210.70

                    SELECT to_char(1210.7, 'L9G999.99');
                    $ 1,210.70

                    SELECT to_char(121, '9 9 9');
                    1 2 1

                    SELECT to_char(121, '00999');
                    00121


                with dates:
                -----------

                Parameter	Explanation
                ----------------------------
                YYYY	4-digit year
                Y,YYY	4-digit year, with comma
                YYY
                YY
                Y	Last 3, 2, or 1 digit(s) of year
                IYYY	4-digit year based on the ISO standard
                IYY
                IY
                I	Last 3, 2, or 1 digit(s) of ISO year
                Q	Quarter of year (1, 2, 3, 4; JAN-MAR = 1).
                MM	Month (01-12; JAN = 01).
                MON	Abbreviated name of month in all uppercase
                Mon	Abbreviated name of month capitalized
                mon	Abbreviated name of month in all lowercase
                MONTH	Name of month in all uppercase, padded with blanks to length of 9 characters
                Month	Name of month capitalized, padded with blanks to length of 9 characters
                month	Name of month in all lowercase, padded with blanks to length of 9 characters
                RM	Month in uppercase Roman numerals
                rm	Month in lowercase Roman numerals
                WW	Week of year (1-53) where week 1 starts on the first day of the year
                W	Week of month (1-5) where week 1 starts on the first day of the month
                IW	Week of year (01-53) based on the ISO standard
                DAY	Name of day in all uppercase, padded with blanks to length of 9 characters
                Day	Name of day capitalized, padded with blanks to length of 9 characters
                day	Name of day in all lowercase, padded with blanks to length of 9 characters
                DY	Abbreviated name of day in all uppercase
                Dy	Abbreviated name of day capitalized
                dy	Abbreviated name of day in all lowercase
                DDD	Day of year (1-366)
                IDDD	Day of year based on ISO year
                DD	Day of month (01-31)
                D	Day of week (1-7, where 1=Sunday, 7=Saturday)
                ID	Day of week based on ISO year (1-7, where 1=Monday, 7=Sunday)
                J	Julian day; the number of days since midnight on November 24, 4714 BC
                HH	Hour of day (01-12)
                HH12	Hour of day (01-12)
                HH24	Hour of day (00-23)
                MI	Minute (00-59)
                SS	Second (00-59)
                MS	Millisecond (000-999)
                US	Microsecond (000000-999999)
                SSSS	Seconds past midnight (0-86399)
                am, AM, pm, or PM	Meridian indicator
                a.m., A.M., p.m., or P.M.	Meridian indicator
                ad, AD, a.d., or A.D	AD indicator
                bc, BC, b.c., or B.C.	BC indicator
                TZ	Name of time zone in uppercase
                tz	Name of time zone in lowercase
                CC	2-digit century


                SELECT to_char(date '2014-04-25', 'YYYY/MM/DD');
                2014/04/25

                SELECT to_char(date '2014-04-25', 'MMDDYY');
                042514

                SELECT to_char(date '2014-04-25', 'Month DD, YYYY');
                April     25, 2014


            23) to_date( string1, format_mask )

                SELECT to_date('2014/04/25', 'YYYY/MM/DD');
                to_date
                ------------
                2014-04-25

                postgres=# SELECT to_date('033114', 'MMDDYY');
                to_date
                ------------
                2014-03-31

                postgres=# SELECT to_date('February 08, 2014', 'Month DD, YYYY');
                to_date
                ------------
                2014-02-08


            24) to_number( string1, format_mask )

                SELECT to_number('1210.73', '9999.99');
                to_number
                -----------
                1210.73

                postgres=# SELECT to_number('1,210.73', '9G999.99');
                to_number
                -----------
                1210.73

                postgres=# SELECT to_number('$1,210.73', 'L9G999.99');
                to_number
                -----------
                1210.73

                postgres=# SELECT to_number('$1,210.73', 'L9G999');
                to_number
                -----------
                    1210
        ==============================================================================

        CREATE TABLE orders (
            info json NOT NULL
        );

        Insert JSON data:
        -----------------
        INSERT INTO orders (info)
        VALUES('{ "customer": "Srinivas", "items": {"product": "keyboard","qty": 2}}');

        Querying JSON data:
        -------------------
        SELECT info FROM orders;