package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) All() ([]*Member, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT * FROM member ORDER BY mem_id ;`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var allMember []*Member

	for rows.Next() {
		var member Member

		err := rows.Scan(
			&member.ID,
			&member.Name,
			&member.Debut_date,
			&member.Birth_date,
		)

		if err != nil {
			return nil, err
		}

		genQuery :=
			` SELECT mg.member_gen_id , m.mem_id , g.gen_id , g.gen_name FROM 
			(
				(
					member m 
					INNER JOIN member_gen mg ON m.mem_id = mg.mem_id
				) 
				INNER JOIN generation g ON g.gen_id = mg.gen_id
			) 
			WHERE m.mem_id = $1;
		`
		genRows, _ := m.DB.QueryContext(ctx, genQuery, member.ID)

		generation := make(map[int]string)

		for genRows.Next() {
			var memGen MemberGeneration
			err := genRows.Scan(
				&memGen.ID,
				&memGen.Mem_id,
				&memGen.Gen_id,
				&memGen.Generation.Name,
			)

			if err != nil {
				return nil, err
			}
			generation[memGen.Gen_id] = memGen.Generation.Name
		}
		genRows.Close()
		member.MemberGen = generation
		allMember = append(allMember, &member)

	}

	return allMember, nil
}
