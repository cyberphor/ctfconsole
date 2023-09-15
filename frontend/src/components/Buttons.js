import Button from 'react-bootstrap/Button';
import React, { useState, useEffect } from 'react';

export default function ButtonGetPlayers() {
  const [data, setData] = useState(null);

  const fetchData = async () => {
    try {
      const response = await fetch('http://localhost:8443/api/v1/player');
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const jsonData = await response.json();
      setData(jsonData);
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };

  useEffect(() => {
    // This function will run when the component is mounted

    // You can call fetchData here if you want to load data when the component is mounted.

    // Cleanup function (optional)
    return () => {
      // You can perform cleanup tasks here if needed
    };
  }, []); // Empty dependency array means this effect runs once when the component is mounted

  return (
    <div>
      <Button onClick={() => fetchData()}>Players</Button>
      {data && (
        <div>
          <pre>{JSON.stringify(data, null, 2)}</pre>
        </div>
      )}
    </div>
  );
}