import {GetSchedule} from '../wailsjs/go/main/App'
import {useEffect, useState} from "react";

function Schedule(props) {
    const [scheduleList, setScheduleList] = useState([]);

    useEffect(() => {
        GetSchedule().then((_scheduleList) => {
            setScheduleList(_scheduleList.list);
        })
    }, []);

    return (
        <div>
            <h2>{props.selectedMenuItem}</h2>
            {scheduleList.map((item, index) => (
                <div>
                    <h3>{item.date}</h3>
                    <h3>{item.problems[0].name}</h3>
                    <h3>{item.problems[0].platform}</h3>
                    <h3>{item.problems[0].url}</h3>
                </div>
            ))}
            <button onClick={() => {
                props.setSelectedMenuItem(null)
            }}>Go Back
            </button>
        </div>
    );
}

export default Schedule