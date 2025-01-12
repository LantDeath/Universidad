% Solución Laboratorio de Física 1 - UNIR.
% Prof. Dr. Miguel Cabeza.
% Alumn. Liz Ángel Núñez.
clear;
% Parámetros iniciales.
v1 = 2; % velocidad de nave a rescatar.
v2 = 10; % velocidad de nave rescatadora.
tc = 900; % tiempo en comenzar la búsqueda desde la posición reportada.
delta_t = 15; % tiempo en cambiar rumbo.
D = (v1 * v2 * tc) / (v2 - v1); % distancia de interceptación.
ti = D / v1; % tiempo en el que se comenzaría la maniobra angular.

% Nuevos datos.
R =6.371009e6; % Radio Terrestre. 
lat = 38.275646; % Latitud inicial.
long = -0.428773; % Longitud inicial.

% Conversión de latitud y longitud a radianes.
lat_rad = deg2rad(38.275646); % Latitud en radianes.
long_rad =deg2rad(-0.428773); % Longitud en radianes.

% Parte líneal.
t = 0:delta_t:ti; % vector tiempo en maniobra lineal.
r = v1 * t; % Distancia recorrida
theta = 0 * t; % es un vector de la misma dimensión de t y r.


% Convertir trayectoria líneal.
lat_lineal = lat_rad + zeros(size(r));
long_lineal = long_rad + (r / R);

%Conversión a grados.
lat_lineal_g = rad2deg(lat_lineal);
long_lineal_g = rad2deg(long_lineal);

pos_lineal = [lat_lineal_g(:), long_lineal_g(:)];

% Cálculo de la trayectoria angular.
theta_max = 2*pi; % ángulo máximo
r_max = D * exp(theta_max / sqrt ((v2 / v1)^2 - 1)); % distancia máxima desde el origen
t_max = r_max / v1; % tiempo máximo de busqueda

% Cálculo del recorrido angular.
t = ti:delta_t:t_max; % vector tiempo recorrido en maniobra angular.
theta_angular = sqrt((v2 / v1)^2 - 1) * log(1 + (v1 * (t - ti)) / D);
r_angular = D * exp(theta_angular / sqrt((v2 / v1)^2 - 1));

% Convertir la trayectoria angular.
latitudes_angular = lat_rad + (r_angular .* sin(theta_angular)) / R;
longitudes_angular = long_rad +(r_angular .* cos(theta_angular)) / R;

% Conversión a grados.
lat_G=rad2deg(latitudes_angular);
long_G=rad2deg(longitudes_angular);


% CREO MATRIZ POSICION ANGULAR.
pos_angular = [lat_G', long_G'];

% Unión de las matrices lineal y angular.
posiciones = [pos_lineal; pos_angular];

csvwrite("posiciones.csv", posiciones);

