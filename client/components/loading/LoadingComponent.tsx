import React from "react";
import { BallTriangle } from "react-loader-spinner";

function LoadingComponent() {
  return (
    <BallTriangle
      height={100}
      width={100}
      radius={5}
      color="#4fa94d"
      ariaLabel="ball-triangle-loading"
      visible={true}
    />
  );
}

export default LoadingComponent;
