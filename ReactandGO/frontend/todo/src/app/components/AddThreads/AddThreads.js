// app/threads/AddThread.js (Client Component)
'use client';

import { useState } from 'react';
import axios from 'axios';

export default function AddThread() {
  const [newThread, setNewThread] = useState([]);
  const[newThreadTitle,setNewThreadTitle]= useState("")
  
  const handleAddThread = async() => {
    setNewThread((prevNewThread)=>[...prevNewThread,{title:newThreadTitle,completed:false},])
    setNewThreadTitle('');
    const response=await axios.post("http://localhost:9000/todo",{title:newThreadTitle,completed:false},{
            headers: {
              'Content-Type': 'application/json'  // Make sure to specify the Content-Type
            }
        }
    )
    console.log(response)

  };

  return (

    <div>
        <div>
            {newThread.map((thread,index)=>(
                <div key={index}>{thread.title}</div>
            ))}
        </div>
      <input
        name='title'
        type="text"
        value={newThreadTitle}
        onChange={(e) =>setNewThreadTitle(e.target.value)}  
        placeholder="Add a new thread"
      />
      <button onClick={handleAddThread}>Add Thread</button>
    </div>
  );
}

