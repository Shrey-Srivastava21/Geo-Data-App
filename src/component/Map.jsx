import React, { useState } from 'react';
import { MapContainer, TileLayer, FeatureGroup, GeoJSON } from 'react-leaflet';
import { EditControl } from 'react-leaflet-draw';
import 'leaflet-draw/dist/leaflet.draw.css';
import "leaflet/dist/leaflet.css";

const Map = ({ geojsonData }) => {
  const [drawnShapes, setDrawnShapes] = useState([]);

  const handleDrawCreated = (event) => {
    const { layerType, layer } = event;
    const shapeData = {};
    if (layerType === 'polygon') {
      shapeData.type = 'Polygon';
      shapeData.coordinates = layer.getLatLngs().map(latlng => [latlng.lng, latlng.lat]);
    } else if (layerType === 'circle') {
      shapeData.type = 'Circle';
      shapeData.center = [layer.getLatLng().lng, layer.getLatLng().lat];
      shapeData.radius = layer.getRadius();
    }
    setDrawnShapes([...drawnShapes, shapeData]);
  };

  const handleDownloadGeoJSON = () => {
    const geoJSONData = {
      type: 'FeatureCollection',
      features: drawnShapes.map(shape => ({
        type: 'Feature',
        geometry: {
          type: shape.type === 'Polygon' ? 'Polygon' : 'Point',
          coordinates: shape.type === 'Polygon' ? [shape.coordinates] : shape.center,
        },
        properties: shape.type === 'Polygon' ? {} : { radius: shape.radius },
      })),
    };
    const blob = new Blob([JSON.stringify(geoJSONData)], { type: 'application/json' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'drawn_shapes.geojson';
    a.click();
  };

  return (
    <div style={{ width: '90%', margin: 'auto', borderRadius: '5px', textAlign: 'center' }}>
      <h2 style={{ margin: '20px 0' }}>​𝗠𝗮𝗽 𝗩𝗶𝗲𝘄𝗲𝗿</h2>
      <MapContainer center={[51.505, -0.09]} zoom={13} style={{ height: '80vh', width: '100%', borderRadius: '10px', boxShadow: '0 2px 10px rgba(0, 0, 0, 0.2)' }}>
        <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
        {geojsonData && <GeoJSON data={JSON.parse(geojsonData)} />}
        <FeatureGroup>
          <EditControl
            position="topright"
            onCreated={handleDrawCreated}
            draw={{
              rectangle: false,
              marker: false,
            }}
          />
        </FeatureGroup>
      </MapContainer>
      <button 
        onClick={handleDownloadGeoJSON} 
        style={{ 
          marginTop: '20px', 
          padding: '10px 20px', 
          backgroundColor: '#ff4d4d', 
          color: '#fff', 
          border: 'none', 
          borderRadius: '5px', 
          cursor: 'pointer',
          transition: 'background-color 0.3s',
        }}
        onMouseEnter={(e) => e.target.style.backgroundColor = '#ff1a1a'}
        onMouseLeave={(e) => e.target.style.backgroundColor = '#ff4d4d'}
      >
        Download GeoJSON
      </button>
        <p></p>
    </div>
  );
};

export default Map;
