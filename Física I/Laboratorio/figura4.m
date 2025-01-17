% Solución Laboratorio de Física 1 - UNIR
% Prof. Dr. Miguel Cabeza
% Alumn. Liz Ángel Núñez 

clear; % Elimina las variables, quitandolas de la memoria del sistema.
clf; % "reset" elimina los elementos secundarios y restablece las propiedas
% en sus valores predeterminados.

v1=2; % velocidad de nave a rescatar.
v2=10; % velocidad de nave rescatadora.
tc=900; % tiempo en comenzar la búsqueda desde la posición reportada.
delta_t =15; % tiempo en cambiar rumbo.
D=(v1*v2*tc)/(v2-v1); % distancia de interceptación.
ti=D/v1; % tiempo en el que se comezaría la maniobra angular.
 
% generación de la gráfica del recorrido lineal.

t= [0,delta_t,ti]; % vector tiempo en maniobra lineal.
r=v1*t; % r es un vector, ya que t lo es
% theta=0; no sería un vector, ya que t lo es.
theta=0*t; % es un vector de la misma dimensión de t y r.
polarplot(theta,r) % generamos la gráfica.

hold on; % mantenemos el gráfico para añadir el recorrido.

% cálculo del tiempo máximo de búsqueda.

theta1= 2*pi; % ángulo máximo
r = (v1*v2*tc)/(v2-v1)*exp(theta1/sqrt((v2/v1)^2-1)); % distancia máxima desde el origen
tmax=r/v1; % tiempo máximo de busqueda

% generación de la gráfica del recorrido angular.

t = ti:delta_t:tmax; % vector tiempo recorrido en maniobra angular.
theta = sqrt((v2 / v1)^2 - 1) * log(1 + (v1 * (t - ti)) / D);
r = D * exp(theta/sqrt((v2/v1)^2 - 1));
polarplot(theta, r);

title ("Recorrido líneal y angular")
