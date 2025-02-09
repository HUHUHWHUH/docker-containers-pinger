import React, { useState, useEffect } from "react";
import { getPingResults } from "../services/pingService";
import PingResultsTable from "./pingResultsTable";
import axios from "axios";

const PingResults = () => {
    const [pingResults, setPingResults] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    // Функция для получения данных с бэкенда
    const fetchPingResults = async () => {
        try {
            const data = await getPingResults();
            setPingResults(data);
            setError(null);
        } catch (err) {
            setError("Ошибка при загрузке данных");
        } finally {
            setLoading(false);
        }
    };

    // Функция для удаления контейнеров со статусом "shutdown"
    const handleDeleteShutdownContainers = async () => {
        try {
            const response = await axios.delete("http://localhost:9090/containers");
            if (response.status === 200) {
                alert("Контейнеры со статусом 'shutdown' успешно удалены");
                fetchPingResults();
            }
        } catch (err) {
            alert("Не удалось удалить контейнеры");
            console.error("Ошибка при удалении контейнеров:", err);
        }
    };

    // useEffect для загрузки данных и их обновления каждые 10 секунд
    useEffect(() => {
        fetchPingResults();
        const interval = setInterval(() => { fetchPingResults(); }, 10000);
        return () => clearInterval(interval);
    }, []);

    if (loading) {
        return <div className="container mt-4"><h3>Загрузка...</h3></div>;
    }

    if (error) {
        return <div className="container mt-4"><h3>{error}</h3></div>;
    }

    if (pingResults.length === 0) {
        return <div className="container mt-4"><h3>Нет данных для отображения</h3></div>;
    }

    return (
        <div className="container mt-4">
            <h2 className="mb-4">Результаты пинга контейнеров</h2>
            <button className="btn btn-danger mb-4" onClick={handleDeleteShutdownContainers}>
                Удалить все контейнеры со статусом "shutdown"
            </button>
            <PingResultsTable results={pingResults} />
        </div>
    );
};

export default PingResults;
