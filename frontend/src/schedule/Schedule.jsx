import React, {useEffect, useState} from 'react';
import {GetSchedule} from '../../wailsjs/go/main/App';
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
                    <span className='date'>{item.date}</span>
                    {item.problems.map((problem, pi) => (
                        <div onClick={() => {
                            window.open(problem.url, '_blank');
                        }}>
                            {pi + 1 !== item.problems.length ? (
                                <div className='problem'>
                                    <span className='problemDetail title'>{pi + 1}. {problem.name}</span>
                                    <span className='problemDetail platform'>{problem.platform}</span>
                                </div>) : (
                                <div className='problem problemLast'>
                                    <span className='problemDetail title'>{pi + 1}. {problem.name}</span>
                                    <span className='problemDetail platform'>{problem.platform}</span>
                                </div>)
                            }
                        </div>
                    ))}
                </div>
            ))}
            <button className='goBackButton' onClick={() => {
                props.setSelectedMenuItem(null)
            }}>Go Back
            </button>
        </div>
    );
}

export default Schedule;
