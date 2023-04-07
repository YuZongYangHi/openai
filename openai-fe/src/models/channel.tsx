import {useState} from "react";

export default () => {
  const [onRouteChange, setOnRouteChange] = useState(0)
  return {onRouteChange, setOnRouteChange}
}
