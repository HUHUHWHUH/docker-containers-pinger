import React from "react";

const PingResultsTable = ({ results }) => {
    return (
        <div className="table-responsive">
            <table className="table table-striped table-hover">
                <thead className="thead-dark">
                <tr>
                    <th>Название контейнера</th>
                    <th>ID контейнера</th>
                    <th>IP-адрес</th>
                    <th>Время пинга</th>
                    <th>Последний успешный пинг</th>
                    <th>Статус</th>
                </tr>
                </thead>
                <tbody>
                {results.map((result) => (
                    <tr key={result.id}>
                        <td>{result.container_name}</td>
                        {<td>{result.container_id}</td>}
                        <td>{result.ip}</td>
                        <td>{new Date(result.pingTime).toLocaleString()}</td>
                        <td>{new Date(result.lastSuccessfulPingTryTime).toLocaleString()}</td>
                        <td
                            className={
                                result.status.toLowerCase() === "shutdown" ? "text-danger" : "text-success"
                            }
                        >
                            {result.status}
                        </td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

export default PingResultsTable;
