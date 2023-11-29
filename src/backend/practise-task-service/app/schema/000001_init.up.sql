CREATE TABLE practice_info (
  Id SERIAL PRIMARY KEY,
  Relative_path varchar,
  Author varchar,
  Title varchar,
  Theme varchar,
  Academic_Subject varchar,
  Created_at timestamp,
  Deleted_at timestamp default null
);
CREATE TABLE access_groups (
    Group_Id SERIAL PRIMARY KEY,
    Group_Name varchar
);
CREATE TABLE practice_access (
                            Practice_Id INT,
                            Group_Id INT,
                            PRIMARY KEY (Practice_Id, Group_Id),
                            FOREIGN KEY (Practice_Id) REFERENCES practice_info(Id),
                            FOREIGN KEY (Group_Id) REFERENCES access_groups(Group_Id)
);
CREATE INDEX idx_title_tsvector ON practice_info USING gin(to_tsvector('russian', title));
CREATE INDEX idx_subject_tsvector ON practice_info USING gin(to_tsvector('russian', academic_subject))

