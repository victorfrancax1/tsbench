package query

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

func DoSelectQuery(query Query) time.Duration {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Unable to connect to database:", err)
	}
	defer conn.Close(context.Background())

	queryString := fmt.Sprintf(`
		SELECT time_bucket('1 minute', ts) as one_min,
			MAX(usage) as max_usage,
			MIN(usage) as min_usage
		FROM cpu_usage
		WHERE host = '%s'
			AND ts >= '%s'
			AND ts <= '%s'
		GROUP BY one_min`,
		query.Host,
		query.StartTime,
		query.EndTime)

	start := time.Now()

	rows, err := conn.Query(context.Background(), queryString)
	if err != nil {
		fmt.Println("QueryRow failed:", err)
	}

	elapsed := time.Since(start)

	rows.Close()
	return elapsed
}
