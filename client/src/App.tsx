import { ModeToggle } from "./components/ModeToggle";
import { useQuery } from "@tanstack/react-query";
import TodoCard from "./components/TodoCard";

export interface Todo {
  ID: string;
  Title: string;
  Body: string;
  IsCompleted: boolean;
}

const getTodos = async () => {
  try {
    const response = await fetch("http://localhost:3000/todos");
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }
    const data: Todo[] = await response.json();
    return data;
  } catch (error) {
    throw new Error(`Fetch failed: ${error}`);
  }
};

function App() {
  const { data, isError, isFetching } = useQuery({
    queryKey: ["todos"],
    queryFn: getTodos,
  });

  return (
    <div>
      <ModeToggle />
      {isFetching && <p>Data fetching....</p>}
      {data && data.map((e) => <TodoCard todo={e} />)}
    </div>
  );
}

export default App;
