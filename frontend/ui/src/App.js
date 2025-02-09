import React, { useEffect } from "react";
import PingResults from "./components/pingResults";

function App() {
    useEffect(() => {
        const interval = setInterval(() => {
            window.location.reload();
        }, 10000);

        return () => clearInterval(interval);
    }, []);

    return (
        <div className="App">
            <div className="container mt-4">
                <h1 className="mb-4">Пингер докер контейнеров запущен, цикл опроса контейнеров - 10 секунд</h1>
                <PingResults/>
            </div>
        </div>
    );
}

export default App;
