import React from "react";
import { Todo } from "../model";
import SingleTodo from "./SingleTodo";
import "./styles.css";

interface Props {
  todos: Todo[];
  setTodos: React.Dispatch<React.SetStateAction<Todo[]>>;
}

const TodoList: React.FC<Props> = ({ todos, setTodos }) => {
  return (
    <div className="todos">
      {todos.map((t) => (
        <SingleTodo todo={t} key={t.id} todos={todos} setTodos={setTodos} />
      ))}
    </div>
  );
};

export default TodoList;
