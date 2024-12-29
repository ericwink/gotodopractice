import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import type { Todo } from "@/App";
import { Button } from "./ui/button";
import { useMutation, useQueryClient } from "@tanstack/react-query";

interface Props {
  todo: Todo;
}

const deleteTodo = async (id: string) => {
  try {
    const response = await fetch("http://localhost:3000/todos", {
      method: "Delete",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ID: id }), // Send the id in the request body
    });
    if (!response.ok) {
      throw new Error(`Failed to delete todo: ${response.status}`);
    }
    const data: Todo[] = await response.json();
    return data;
  } catch (error) {
    throw new Error(`Fetch failed: ${error}`);
  }
};

const TodoCard = ({ todo }: Props) => {
  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: deleteTodo,
    onSuccess: () => {
      // Invalidate and refetch
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });

  return (
    <Card>
      <CardHeader>
        <CardTitle>{todo.Title}</CardTitle>
      </CardHeader>
      <CardContent>
        <p>{todo.Body}</p>
        <p>{todo.IsCompleted === true ? "Completed" : "Not completed"}</p>
      </CardContent>
      <CardFooter>
        <Button onClick={() => mutation.mutate(todo.ID)}>Delete Task</Button>
      </CardFooter>
    </Card>
  );
};

export default TodoCard;
