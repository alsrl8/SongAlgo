import {GetSchedule} from '../wailsjs/go/main/App'
import {useEffect, useState} from "react";

function Schedule({selectedMenuItem, setSelectedMenuItem}) {
    const [schedule, setSchedule] = useState(null);

    useEffect(() => {
        GetSchedule().then((_schedule) => {
            setSchedule(_schedule);
        })
    }, []);

    return (
        <div>
            <h2>{selectedMenuItem}</h2>
            <button onClick={setSelectedMenuItem(null)}>Go Back</button>
        </div>
    );
}

export default Schedule