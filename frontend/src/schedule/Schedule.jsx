import React, { useEffect, useState } from 'react';
import { GetSchedule } from '../../wailsjs/go/main/App';
import './Schedule.css'

function Schedule(props) {
    const [scheduleList, setScheduleList] = useState([]);

    useEffect(() => {
        GetSchedule().then((_scheduleList) => {
            setScheduleList(_scheduleList.list);
        });
    }, []);

    return (
        <div className='scheduleContainer'>
            <h2>{props.selectedMenuItem}</h2>
            {scheduleList.map((item, index) => (
                <div key={index} className='scheduleCard'>
                    <div className='date'>{item.date}</div>
                    <div className='problemDetails'>
                        <h4>{item.problems[0].name}</h4>
                        <p>Platform: {item.problems[0].platform}</p>
                        <p>URL: <a href={item.problems[0].url} target="_blank" rel="noopener noreferrer">{item.problems[0].url}</a></p>
                    </div>
                </div>
            ))}
            <button className='goBackButton' onClick={() => { props.setSelectedMenuItem(null) }}>Go Back</button>
        </div>
    );
}

export default Schedule;
