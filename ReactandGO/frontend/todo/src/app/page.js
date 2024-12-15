// app/threads/page.js (Server Component)
import AddThread from "./components/AddThreads/AddThreads";
import axios from "axios";

async function fetchThreads() {
  const response = await fetch("http://localhost:9000/todo",{next:{revalidate:60}})
  
  const data=await response.json()
  console.log(data)
  return data.todos
}

export default async function ThreadsPage() {
  let threads = await fetchThreads();
  if(!Array.isArray(threads)){
    threads=[]
  }

  return (
    <div>
      <h1>Todos</h1>
      <ul>
        {threads.map((thread) => (
          <li key={thread.todoid}>{thread.title}</li>
        ))}
      </ul>
    </div>
  );
}

