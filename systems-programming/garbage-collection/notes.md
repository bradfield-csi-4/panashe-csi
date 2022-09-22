Running with garbage collector tracing definitely slows it down.
GODEBUG=gctrace=1 prints garbage collector events, and summarizes the amount of garbage collected and the length of the pause.


```
2022/09/21 09:16:09.647485 service.go:64: Listening on: 0.0.0.0:5000
gc 1 @32.385s 0%: 0.034+1.1+0.004 ms clock, 0.27+0.23/1.9/3.6+0.037 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
2022/09/21 09:16:42.049950 rss.go:109: reloaded cache http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml
gc 2 @32.408s 0%: 0.099+1.8+0.037 ms clock, 0.79+1.7/3.0/0.85+0.30 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
2022/09/21 09:16:42.081662 rss.go:109: reloaded cache http://rss.nytimes.com/services/xml/rss/nyt/US.xml
2022/09/21 09:16:42.107530 rss.go:109: reloaded cache http://rss.nytimes.com/services/xml/rss/nyt/Politics.xml
gc 3 @32.467s 0%: 0.072+0.99+0.079 ms clock, 0.57+0.95/1.4/0.88+0.63 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
2022/09/21 09:16:42.132469 rss.go:109: reloaded cache http://rss.nytimes.com/services/xml/rss/nyt/Business.xml
2022/09/21 09:16:42.135706 rss.go:109: reloaded cache http://feeds.bbci.co.uk/news/rss.xml
gc 4 @32.493s 0%: 0.081+1.0+0.026 ms clock, 0.65+0.89/1.4/1.2+0.21 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
2022/09/21 09:16:42.223143 rss.go:109: reloaded cache http://feeds.bbci.co.uk/news/world/rss.xml
2022/09/21 09:16:42.262508 rss.go:109: reloaded cache http://rss.cnn.com/rss/cnn_topstories.rss
gc 5 @32.620s 0%: 0.058+0.86+0.002 ms clock, 0.46+1.4/1.2/0.088+0.020 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
2022/09/21 09:16:42.360831 rss.go:109: reloaded cache http://feeds.bbci.co.uk/news/politics/rss.xml
gc 6 @32.720s 0%: 0.10+0.81+0.099 ms clock, 0.87+1.3/1.2/0.76+0.79 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
2022/09/21 09:16:42.442081 rss.go:109: reloaded cache http://rss.cnn.com/rss/cnn_world.rss
2022/09/21 09:16:42.709142 rss.go:109: reloaded cache http://rss.cnn.com/rss/cnn_us.rss
gc 7 @33.068s 0%: 0.11+1.4+0.006 ms clock, 0.88+1.5/2.6/1.9+0.051 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
2022/09/21 09:16:42.890571 rss.go:109: reloaded cache http://rss.cnn.com/rss/cnn_allpolitics.rss
2022/09/21 09:16:42.995440 rss.go:109: reloaded cache http://feeds.bbci.co.uk/news/world/us_and_canada/rss.xml
gc 8 @33.356s 0%: 0.19+2.1+0.037 ms clock, 1.5+1.2/2.4/0+0.29 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 9 @33.361s 0%: 0.21+1.5+0.16 ms clock, 1.7+2.8/1.7/0+1.3 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 10 @33.364s 0%: 0.41+1.6+0.16 ms clock, 3.2+3.4/2.0/0+1.2 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 11 @33.369s 0%: 0.24+3.4+0.21 ms clock, 1.9+3.4/1.9/0+1.7 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 12 @33.376s 0%: 0.33+1.2+0.049 ms clock, 2.6+2.3/1.5/0+0.39 ms cpu, 6->7->2 MB, 7 MB goal, 8 P
gc 13 @33.379s 0%: 0.24+1.3+0.24 ms clock, 1.9+2.4/1.5/0+1.9 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 14 @33.383s 0%: 0.16+1.2+0.032 ms clock, 1.3+2.4/1.4/0+0.26 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 15 @33.386s 0%: 0.15+2.0+0.045 ms clock, 1.2+2.8/2.3/0+0.36 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 16 @33.391s 0%: 0.13+1.1+0.003 ms clock, 1.1+2.0/1.3/0+0.029 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 17 @33.394s 0%: 0.18+1.6+0.052 ms clock, 1.5+2.7/1.8/0+0.42 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 18 @33.398s 0%: 0.20+1.4+0.094 ms clock, 1.6+2.2/1.8/0+0.75 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 19 @33.402s 0%: 0.15+1.0+0.095 ms clock, 1.2+2.1/1.1/0+0.76 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 20 @33.404s 0%: 0.090+2.2+0.086 ms clock, 0.72+2.2/2.1/0+0.69 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 21 @33.410s 0%: 0.14+1.1+0.041 ms clock, 1.1+1.5/1.5/0+0.33 ms cpu, 6->7->2 MB, 7 MB goal, 8 P
gc 22 @33.413s 0%: 0.18+1.7+0.057 ms clock, 1.5+1.6/1.2/0+0.45 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 23 @33.417s 0%: 0.19+1.1+0.050 ms clock, 1.5+1.8/1.4/0+0.40 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 24 @33.420s 0%: 0.084+1.6+0.046 ms clock, 0.67+3.5/1.5/0+0.37 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 25 @33.424s 0%: 0.22+1.3+0.018 ms clock, 1.7+1.4/1.6/0+0.14 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 26 @33.427s 0%: 0.15+1.0+0.063 ms clock, 1.2+1.2/1.3/0+0.51 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 27 @33.430s 0%: 0.098+0.83+0.042 ms clock, 0.78+1.5/1.0/0+0.33 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 28 @33.432s 0%: 0.13+1.6+0.089 ms clock, 1.0+1.5/1.2/0+0.71 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 29 @33.436s 0%: 0.13+1.5+0.10 ms clock, 1.1+3.0/0.80/0+0.86 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 30 @33.439s 0%: 0.14+1.1+0.090 ms clock, 1.1+1.1/1.5/0+0.72 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 31 @33.443s 0%: 0.15+1.0+0.004 ms clock, 1.2+2.4/1.1/0+0.035 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 32 @33.446s 0%: 0.091+1.2+0.013 ms clock, 0.72+2.4/1.4/0+0.11 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 33 @33.449s 0%: 0.058+0.74+0.023 ms clock, 0.47+1.4/0.93/0+0.18 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 34 @33.451s 0%: 0.038+2.3+0.063 ms clock, 0.30+0.95/0.99/0+0.51 ms cpu, 5->7->4 MB, 6 MB goal, 8 P
gc 35 @33.456s 0%: 0.094+1.2+0.003 ms clock, 0.75+0.67/1.7/0+0.024 ms cpu, 8->8->2 MB, 9 MB goal, 8 P
gc 36 @33.459s 0%: 0.16+1.1+0.048 ms clock, 1.2+0.81/1.0/0+0.38 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 37 @33.462s 0%: 0.11+0.94+0.006 ms clock, 0.94+1.6/1.0/0+0.053 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 38 @33.464s 0%: 0.062+1.5+0.012 ms clock, 0.49+3.0/0.82/0+0.10 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 39 @33.468s 0%: 0.056+1.5+0.020 ms clock, 0.44+2.4/1.7/0+0.16 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 40 @33.470s 0%: 0.094+2.3+0.085 ms clock, 0.75+1.2/1.6/0+0.68 ms cpu, 5->7->4 MB, 7 MB goal, 8 P
gc 41 @33.474s 0%: 0.12+2.1+0.10 ms clock, 1.0+4.4/1.6/0+0.82 ms cpu, 6->8->3 MB, 8 MB goal, 8 P
gc 42 @33.478s 0%: 0.088+1.1+0.026 ms clock, 0.70+1.5/1.2/0+0.21 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 43 @33.480s 0%: 0.099+0.71+0.010 ms clock, 0.79+0.78/0.98/0+0.087 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 44 @33.482s 0%: 0.064+0.96+0.017 ms clock, 0.51+0.65/1.2/0+0.13 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 45 @33.484s 0%: 0.075+1.1+0.020 ms clock, 0.60+0.77/1.3/0+0.16 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 46 @33.487s 0%: 0.11+1.0+0.043 ms clock, 0.92+1.4/1.3/0+0.35 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 47 @33.490s 0%: 0.12+1.0+0.026 ms clock, 0.98+1.5/1.0/0+0.21 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 48 @33.493s 0%: 0.11+0.97+0.062 ms clock, 0.94+1.5/1.3/0+0.49 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 49 @33.496s 0%: 0.12+1.0+0.013 ms clock, 1.0+1.1/1.3/0+0.10 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 50 @33.498s 0%: 0.072+0.91+0.090 ms clock, 0.57+1.4/0.88/0+0.72 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 51 @33.501s 0%: 0.10+0.87+0.034 ms clock, 0.85+1.2/1.0/0+0.27 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 52 @33.503s 0%: 0.10+0.80+0.046 ms clock, 0.87+0.96/0.79/0+0.37 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 53 @33.505s 0%: 0.12+1.3+0.008 ms clock, 0.96+1.5/1.5/0+0.071 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 54 @33.509s 0%: 0.12+1.2+0.078 ms clock, 0.97+1.4/1.3/0+0.62 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 55 @33.512s 0%: 0.084+1.5+0.003 ms clock, 0.67+1.8/2.0/0+0.027 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 56 @33.515s 0%: 0.16+0.83+0.096 ms clock, 1.3+1.3/1.0/0+0.77 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 57 @33.518s 0%: 0.13+1.0+0.079 ms clock, 1.1+1.3/1.3/0+0.63 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 58 @33.521s 0%: 0.11+3.1+0.017 ms clock, 0.92+1.8/2.0/0+0.14 ms cpu, 6->8->4 MB, 7 MB goal, 8 P
gc 59 @33.527s 0%: 0.10+1.9+0.056 ms clock, 0.84+0.66/1.6/0+0.45 ms cpu, 8->9->3 MB, 9 MB goal, 8 P
gc 60 @33.531s 0%: 0.12+1.1+0.018 ms clock, 0.97+1.8/1.3/0+0.14 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 61 @33.533s 0%: 0.11+1.2+0.011 ms clock, 0.89+2.0/1.6/0+0.094 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 62 @33.536s 0%: 0.14+0.96+0.003 ms clock, 1.1+1.5/1.0/0+0.027 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 63 @33.538s 0%: 0.11+1.0+0.028 ms clock, 0.93+1.5/1.2/0+0.22 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 64 @33.541s 0%: 0.16+1.5+0.038 ms clock, 1.2+2.1/1.6/0+0.31 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 65 @33.544s 0%: 0.099+1.4+0.048 ms clock, 0.79+1.2/1.3/0+0.39 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 66 @33.547s 0%: 0.15+1.8+0.059 ms clock, 1.2+1.7/1.1/0+0.47 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 67 @33.551s 0%: 0.11+1.2+0.051 ms clock, 0.92+2.0/1.4/0+0.41 ms cpu, 7->7->3 MB, 8 MB goal, 8 P
gc 68 @33.554s 0%: 0.11+0.94+0.072 ms clock, 0.93+1.4/1.1/0+0.57 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 69 @33.556s 0%: 0.12+1.1+0.084 ms clock, 0.99+1.4/1.1/0+0.67 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 70 @33.559s 0%: 0.11+0.96+0.040 ms clock, 0.89+1.2/0.98/0+0.32 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 71 @33.561s 0%: 0.099+1.6+0.012 ms clock, 0.79+4.1/1.0/0+0.099 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 72 @33.565s 0%: 0.13+1.0+0.057 ms clock, 1.0+1.7/1.2/0+0.45 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 73 @33.568s 0%: 0.082+3.8+0.006 ms clock, 0.65+3.9/3.0/0+0.048 ms cpu, 6->8->4 MB, 7 MB goal, 8 P
gc 74 @33.575s 0%: 0.062+1.2+0.089 ms clock, 0.49+1.8/0.77/0+0.71 ms cpu, 8->8->2 MB, 9 MB goal, 8 P
gc 75 @33.578s 0%: 0.13+2.2+0.096 ms clock, 1.1+4.6/0/0+0.77 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 76 @33.581s 0%: 0.099+1.7+0.020 ms clock, 0.79+3.3/0.99/0+0.16 ms cpu, 5->6->3 MB, 7 MB goal, 8 P
gc 77 @33.584s 0%: 0.10+0.95+0.10 ms clock, 0.83+1.3/0.88/0+0.83 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 78 @33.586s 0%: 0.041+2.4+0.016 ms clock, 0.33+0.52/2.8/0+0.12 ms cpu, 5->6->4 MB, 6 MB goal, 8 P
gc 79 @33.591s 0%: 0.089+0.96+0.085 ms clock, 0.71+0.71/1.1/0+0.68 ms cpu, 6->7->2 MB, 8 MB goal, 8 P
gc 80 @33.593s 0%: 0.082+1.2+0.017 ms clock, 0.66+1.5/1.5/0+0.13 ms cpu, 4->5->3 MB, 5 MB goal, 8 P
gc 81 @33.596s 0%: 0.084+1.1+0.062 ms clock, 0.67+0.59/1.3/0+0.50 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 82 @33.599s 0%: 0.076+1.5+0.095 ms clock, 0.61+0.86/1.7/0+0.76 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 83 @33.602s 0%: 0.085+1.1+0.048 ms clock, 0.68+1.6/1.4/0+0.38 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 84 @33.604s 0%: 0.11+0.85+0.027 ms clock, 0.94+1.6/1.1/0+0.22 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 85 @33.607s 0%: 0.11+3.2+0.012 ms clock, 0.94+2.2/3.7/0+0.098 ms cpu, 5->7->4 MB, 6 MB goal, 8 P
gc 86 @33.613s 0%: 0.086+0.92+0.066 ms clock, 0.68+1.7/0.99/0+0.53 ms cpu, 8->8->2 MB, 9 MB goal, 8 P
gc 87 @33.615s 0%: 0.12+0.91+0.032 ms clock, 0.98+1.6/0.98/0+0.26 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 88 @33.619s 0%: 0.11+1.2+0.048 ms clock, 0.88+1.5/1.4/0+0.38 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 89 @33.622s 0%: 0.069+1.5+0.069 ms clock, 0.55+1.3/1.4/0+0.55 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 90 @33.626s 0%: 0.087+1.9+0.051 ms clock, 0.70+1.0/1.8/0+0.41 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 91 @33.630s 0%: 0.096+0.97+0.017 ms clock, 0.77+1.0/1.1/0+0.13 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 92 @33.632s 0%: 0.097+1.0+0.098 ms clock, 0.78+1.5/0.97/0+0.78 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 93 @33.635s 0%: 0.086+1.2+0.11 ms clock, 0.69+1.6/1.1/0+0.88 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 94 @33.638s 0%: 0.068+2.2+0.002 ms clock, 0.55+4.1/1.3/0+0.017 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 95 @33.643s 0%: 0.050+1.7+0.044 ms clock, 0.40+0.73/2.0/0+0.35 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 96 @33.647s 0%: 0.049+1.1+0.098 ms clock, 0.39+1.2/1.2/0+0.78 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 97 @33.650s 0%: 0.097+1.1+0.075 ms clock, 0.77+1.8/1.3/0+0.60 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 98 @33.652s 0%: 0.071+1.2+0.075 ms clock, 0.57+1.1/1.0/0+0.60 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 99 @33.656s 0%: 0.060+1.8+0.10 ms clock, 0.48+0.82/1.3/0+0.82 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 100 @33.661s 0%: 0.40+1.6+0.11 ms clock, 3.2+1.2/1.5/0+0.89 ms cpu, 6->8->3 MB, 7 MB goal, 8 P
gc 101 @33.665s 0%: 0.070+1.3+0.11 ms clock, 0.56+3.7/0/0+0.94 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 102 @33.669s 0%: 0.076+1.2+0.036 ms clock, 0.61+0.35/1.5/0+0.29 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 103 @33.672s 0%: 0.048+2.3+0.10 ms clock, 0.38+1.7/1.6/0+0.83 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 104 @33.677s 0%: 0.060+1.6+0.027 ms clock, 0.48+2.1/0.86/0+0.21 ms cpu, 7->7->3 MB, 8 MB goal, 8 P
gc 105 @33.680s 0%: 0.096+1.2+0.045 ms clock, 0.76+1.3/1.6/0+0.36 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 106 @33.683s 0%: 0.083+1.4+0.11 ms clock, 0.66+1.6/1.4/0+0.94 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 107 @33.686s 0%: 0.096+1.0+0.012 ms clock, 0.76+1.2/1.1/0+0.097 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 108 @33.689s 0%: 0.12+2.3+0.11 ms clock, 0.96+1.1/0.90/0+0.94 ms cpu, 5->7->3 MB, 6 MB goal, 8 P
gc 109 @33.694s 0%: 0.11+1.2+0.11 ms clock, 0.92+3.8/0.94/0+0.90 ms cpu, 7->7->2 MB, 8 MB goal, 8 P
gc 110 @33.697s 0%: 0.080+2.1+0.020 ms clock, 0.64+4.5/2.1/0+0.16 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 111 @33.701s 0%: 0.046+1.8+0.041 ms clock, 0.36+0.44/1.5/0+0.32 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 112 @33.704s 0%: 0.11+0.95+0.059 ms clock, 0.89+1.5/0.97/0+0.47 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 113 @33.706s 0%: 0.17+1.1+0.034 ms clock, 1.4+2.6/1.6/0+0.27 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 114 @33.710s 0%: 0.039+2.2+0.014 ms clock, 0.31+1.2/3.4/0+0.11 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 115 @33.714s 0%: 0.11+0.92+0.020 ms clock, 0.92+1.6/1.2/0+0.16 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 116 @33.716s 0%: 0.11+1.1+0.042 ms clock, 0.90+1.9/1.2/0+0.34 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 117 @33.719s 0%: 0.068+1.2+0.062 ms clock, 0.55+1.2/1.3/0+0.50 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 118 @33.722s 0%: 0.071+1.6+0.068 ms clock, 0.57+2.2/1.7/0+0.54 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 119 @33.726s 0%: 0.061+1.8+0.002 ms clock, 0.49+0.86/1.2/0+0.022 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 120 @33.730s 0%: 0.053+1.5+0.022 ms clock, 0.42+2.2/0.71/0+0.18 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 121 @33.734s 0%: 0.097+1.7+0.073 ms clock, 0.77+5.3/0/0+0.58 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 122 @33.737s 0%: 0.040+1.8+0.019 ms clock, 0.32+3.8/0.60/0+0.15 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 123 @33.742s 0%: 0.054+1.4+0.077 ms clock, 0.43+0.46/1.6/0+0.61 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 124 @33.745s 0%: 0.097+3.6+0.025 ms clock, 0.78+4.8/3.8/0+0.20 ms cpu, 5->7->4 MB, 6 MB goal, 8 P
gc 125 @33.750s 0%: 0.046+1.1+0.085 ms clock, 0.37+0.93/1.4/0+0.68 ms cpu, 7->8->3 MB, 9 MB goal, 8 P
gc 126 @33.752s 0%: 0.055+1.7+0.019 ms clock, 0.44+0.67/2.6/0+0.15 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 127 @33.755s 0%: 0.054+1.0+0.020 ms clock, 0.43+0.49/1.2/0+0.16 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 128 @33.758s 0%: 0.070+1.0+0.051 ms clock, 0.56+0.90/1.2/0+0.41 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 129 @33.761s 0%: 0.10+2.3+0.067 ms clock, 0.81+4.1/1.5/0+0.53 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 130 @33.765s 0%: 0.090+1.0+0.017 ms clock, 0.72+1.2/1.0/0+0.13 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 131 @33.768s 0%: 0.047+1.1+0 ms clock, 0.37+0.60/1.3/0+0.007 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 132 @33.771s 0%: 0.067+1.7+0.12 ms clock, 0.54+1.3/1.1/0+1.0 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 133 @33.775s 0%: 0.058+1.6+0.048 ms clock, 0.47+1.6/1.0/0+0.39 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 134 @33.779s 0%: 0.078+1.4+0.011 ms clock, 0.62+1.7/0.84/0+0.089 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 135 @33.782s 0%: 0.050+1.9+0.071 ms clock, 0.40+0.75/0.89/0+0.57 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 136 @33.787s 0%: 0.099+1.0+0.13 ms clock, 0.79+1.9/1.3/0+1.0 ms cpu, 7->7->3 MB, 8 MB goal, 8 P
gc 137 @33.791s 0%: 0.17+1.9+0.024 ms clock, 1.3+1.5/2.8/0+0.19 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 138 @33.796s 0%: 0.053+2.8+0.016 ms clock, 0.43+0.53/1.0/0+0.12 ms cpu, 6->8->4 MB, 7 MB goal, 8 P
gc 139 @33.802s 0%: 0.10+1.5+0.050 ms clock, 0.85+1.8/1.9/0+0.40 ms cpu, 8->8->3 MB, 9 MB goal, 8 P
gc 140 @33.806s 0%: 0.079+3.1+0.086 ms clock, 0.63+1.7/1.4/0+0.69 ms cpu, 6->8->4 MB, 7 MB goal, 8 P
gc 141 @33.812s 0%: 0.064+1.7+0.13 ms clock, 0.51+2.9/1.1/0+1.0 ms cpu, 8->9->3 MB, 9 MB goal, 8 P
gc 142 @33.815s 0%: 0.082+1.1+0.056 ms clock, 0.66+1.0/1.5/0+0.45 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 143 @33.817s 0%: 0.10+1.1+0.060 ms clock, 0.82+1.4/1.6/0+0.48 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 144 @33.819s 0%: 0.091+0.90+0.064 ms clock, 0.73+0.92/0.92/0+0.51 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 145 @33.822s 0%: 0.081+3.0+0.024 ms clock, 0.65+5.6/1.3/0+0.19 ms cpu, 5->7->4 MB, 6 MB goal, 8 P
gc 146 @33.827s 0%: 0.21+1.9+0.093 ms clock, 1.7+2.6/2.0/0+0.74 ms cpu, 6->8->3 MB, 8 MB goal, 8 P
gc 147 @33.831s 0%: 0.10+0.91+0.10 ms clock, 0.82+1.8/1.1/0+0.82 ms cpu, 5->6->3 MB, 7 MB goal, 8 P
gc 148 @33.833s 0%: 0.072+1.0+0.053 ms clock, 0.58+1.7/1.2/0+0.42 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 149 @33.835s 0%: 0.091+1.8+0.022 ms clock, 0.73+2.6/2.1/0+0.18 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 150 @33.838s 0%: 0.088+1.2+0.023 ms clock, 0.71+0.90/1.4/0+0.18 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 151 @33.841s 0%: 0.086+1.0+0.092 ms clock, 0.69+1.5/1.2/0+0.74 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 152 @33.843s 0%: 0.070+1.3+0.096 ms clock, 0.56+1.2/1.1/0+0.77 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 153 @33.847s 0%: 0.12+3.2+0.017 ms clock, 1.0+4.6/3.1/0+0.14 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 154 @33.854s 0%: 0.11+1.3+0.086 ms clock, 0.89+1.5/1.4/0+0.69 ms cpu, 7->8->3 MB, 8 MB goal, 8 P
gc 155 @33.856s 0%: 0.069+1.5+0.062 ms clock, 0.55+0.38/1.9/0+0.50 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 156 @33.859s 0%: 0.073+2.9+0.20 ms clock, 0.58+3.6/3.2/0+1.6 ms cpu, 5->7->4 MB, 6 MB goal, 8 P
gc 157 @33.864s 0%: 0.093+1.0+0.049 ms clock, 0.74+2.1/1.2/0+0.39 ms cpu, 6->7->3 MB, 8 MB goal, 8 P
gc 158 @33.866s 0%: 0.096+1.2+0.046 ms clock, 0.77+1.0/1.3/0+0.36 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 159 @33.868s 0%: 0.084+1.4+0.011 ms clock, 0.67+0.98/1.7/0+0.095 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 160 @33.871s 0%: 0.097+1.7+0.10 ms clock, 0.78+1.2/1.5/0+0.83 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 161 @33.875s 0%: 0.076+3.0+0.016 ms clock, 0.61+0.29/1.2/0+0.13 ms cpu, 6->7->4 MB, 7 MB goal, 8 P
gc 162 @33.882s 0%: 0.032+2.4+0.13 ms clock, 0.25+2.9/1.7/0+1.0 ms cpu, 6->9->3 MB, 8 MB goal, 8 P
gc 163 @33.885s 0%: 0.11+3.3+0.019 ms clock, 0.92+4.6/0/0+0.15 ms cpu, 6->8->5 MB, 7 MB goal, 8 P
gc 164 @33.891s 0%: 0.094+1.7+0.10 ms clock, 0.75+4.4/0/0+0.86 ms cpu, 8->9->3 MB, 10 MB goal, 8 P
gc 165 @33.895s 0%: 0.39+2.9+0.074 ms clock, 3.1+3.8/0.16/0+0.59 ms cpu, 5->7->4 MB, 6 MB goal, 8 P
gc 166 @33.900s 0%: 0.062+1.9+0.018 ms clock, 0.50+2.8/1.0/0+0.14 ms cpu, 6->7->3 MB, 8 MB goal, 8 P
gc 167 @33.903s 0%: 0.10+1.1+0.12 ms clock, 0.85+1.1/1.0/0+0.99 ms cpu, 5->6->3 MB, 7 MB goal, 8 P
gc 168 @33.905s 0%: 0.10+1.5+0.10 ms clock, 0.81+1.1/1.5/0+0.83 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 169 @33.908s 0%: 0.056+1.5+0.096 ms clock, 0.44+0.16/1.8/0+0.76 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 170 @33.912s 0%: 0.099+1.0+0.10 ms clock, 0.79+1.8/1.1/0+0.85 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 171 @33.914s 0%: 0.10+0.82+0.084 ms clock, 0.83+1.1/0.90/0+0.67 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 172 @33.916s 0%: 0.10+1.7+0.017 ms clock, 0.86+3.6/0.84/0+0.14 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 173 @33.919s 0%: 0.084+0.98+0.068 ms clock, 0.67+1.4/0.89/0+0.54 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 174 @33.922s 0%: 0.096+4.1+0.002 ms clock, 0.77+2.1/4.5/0+0.023 ms cpu, 5->7->4 MB, 6 MB goal, 8 P
gc 175 @33.930s 0%: 0.055+1.2+0.059 ms clock, 0.44+0.85/1.0/0+0.47 ms cpu, 8->9->3 MB, 9 MB goal, 8 P
gc 176 @33.932s 0%: 0.064+2.2+0.11 ms clock, 0.51+1.9/1.7/0+0.95 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 177 @33.937s 0%: 0.056+2.3+0.020 ms clock, 0.45+1.1/1.8/0+0.16 ms cpu, 6->8->3 MB, 7 MB goal, 8 P
gc 178 @33.941s 0%: 0.13+1.4+0.31 ms clock, 1.0+2.0/1.3/0+2.4 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 179 @33.944s 0%: 0.12+1.2+0.052 ms clock, 1.0+1.0/1.8/0+0.42 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 180 @33.947s 0%: 0.083+1.0+0.029 ms clock, 0.66+2.1/1.4/0+0.23 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 181 @33.949s 0%: 0.079+1.1+0.058 ms clock, 0.63+1.8/1.3/0+0.46 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 182 @33.952s 0%: 0.11+1.0+0.015 ms clock, 0.95+1.5/0.91/0+0.12 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 183 @33.954s 0%: 0.056+1.3+0.14 ms clock, 0.45+0.67/1.2/0+1.1 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 184 @33.958s 0%: 0.12+1.1+0.11 ms clock, 1.0+1.1/1.5/0+0.89 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 185 @33.961s 0%: 0.12+3.7+0.024 ms clock, 0.96+4.7/1.1/0+0.19 ms cpu, 6->8->5 MB, 7 MB goal, 8 P
gc 186 @33.968s 0%: 0.077+2.0+0.059 ms clock, 0.61+1.7/1.1/0+0.47 ms cpu, 8->9->3 MB, 10 MB goal, 8 P
gc 187 @33.972s 0%: 0.19+1.4+0.063 ms clock, 1.5+4.0/0/0+0.51 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 188 @33.976s 0%: 0.046+1.6+0.093 ms clock, 0.37+3.1/0/0+0.75 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 189 @33.979s 0%: 0.078+1.5+0.056 ms clock, 0.63+4.3/0/0+0.45 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 190 @33.981s 0%: 0.095+1.6+0.015 ms clock, 0.76+2.4/1.7/0+0.12 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 191 @33.985s 0%: 0.062+1.7+0.069 ms clock, 0.50+0.86/1.2/0+0.55 ms cpu, 5->7->3 MB, 7 MB goal, 8 P
gc 192 @33.987s 0%: 0.041+1.8+0.095 ms clock, 0.33+2.0/1.4/0+0.76 ms cpu, 5->6->3 MB, 7 MB goal, 8 P
gc 193 @33.991s 0%: 0.085+3.2+0.017 ms clock, 0.68+1.3/3.7/0+0.14 ms cpu, 5->7->4 MB, 6 MB goal, 8 P
gc 194 @33.996s 0%: 0.086+1.0+0.041 ms clock, 0.68+1.5/1.3/0+0.33 ms cpu, 6->7->3 MB, 8 MB goal, 8 P
gc 195 @33.998s 0%: 0.11+1.1+0.066 ms clock, 0.89+1.2/1.0/0+0.53 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 196 @34.000s 0%: 0.062+0.88+0.065 ms clock, 0.49+0.93/0.99/0+0.52 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 197 @34.002s 0%: 0.072+1.0+0.054 ms clock, 0.58+0.87/1.0/0+0.43 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 198 @34.006s 0%: 0.10+1.5+0.088 ms clock, 0.80+1.8/1.9/0+0.71 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 199 @34.009s 0%: 0.10+1.5+0.002 ms clock, 0.85+1.9/1.7/0+0.023 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 200 @34.013s 0%: 0.10+1.2+0.087 ms clock, 0.86+1.7/1.1/0+0.70 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 201 @34.016s 0%: 0.052+0.93+0.13 ms clock, 0.42+1.2/1.1/0+1.0 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 202 @34.018s 0%: 0.19+0.94+0.019 ms clock, 1.5+1.4/1.1/0+0.15 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 203 @34.022s 0%: 0.053+2.7+0.025 ms clock, 0.42+0.68/3.1/0+0.20 ms cpu, 5->7->3 MB, 6 MB goal, 8 P
gc 204 @34.029s 0%: 0.081+1.2+0.061 ms clock, 0.64+1.0/1.4/0+0.49 ms cpu, 7->8->3 MB, 8 MB goal, 8 P
gc 205 @34.032s 0%: 0.082+0.85+0.016 ms clock, 0.65+1.5/1.1/0+0.13 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 206 @34.034s 0%: 0.051+1.5+0.049 ms clock, 0.40+0.84/2.0/0+0.39 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 207 @34.037s 0%: 0.091+0.89+0.027 ms clock, 0.72+0.61/1.1/0+0.22 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 208 @34.040s 0%: 0.092+2.3+0.017 ms clock, 0.73+2.1/3.9/0+0.14 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 209 @34.045s 0%: 0.093+1.5+0.018 ms clock, 0.75+1.6/1.9/0+0.14 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 210 @34.048s 0%: 0.050+1.0+0.039 ms clock, 0.40+0.92/1.2/0+0.31 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 211 @34.051s 0%: 0.080+1.1+0.14 ms clock, 0.64+1.5/1.1/0+1.1 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 212 @34.054s 0%: 0.072+1.0+0.013 ms clock, 0.57+0.49/1.3/0+0.11 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 213 @34.057s 0%: 0.10+2.7+0.018 ms clock, 0.86+1.5/4.1/0+0.14 ms cpu, 5->7->4 MB, 6 MB goal, 8 P
gc 214 @34.063s 0%: 0.14+1.0+0.070 ms clock, 1.1+1.2/1.3/0+0.56 ms cpu, 8->8->3 MB, 9 MB goal, 8 P
gc 215 @34.065s 0%: 0.11+0.82+0.013 ms clock, 0.90+1.2/1.1/0+0.10 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 216 @34.067s 0%: 0.057+1.1+0.016 ms clock, 0.46+1.2/1.5/0+0.12 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 217 @34.070s 0%: 0.044+0.97+0.073 ms clock, 0.35+0.71/1.2/0+0.58 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 218 @34.073s 0%: 0.22+1.0+0.014 ms clock, 1.7+1.3/1.1/0+0.11 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 219 @34.076s 0%: 0.68+1.6+0.003 ms clock, 5.4+0.29/2.1/0+0.027 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 220 @34.080s 0%: 0.10+1.0+0.008 ms clock, 0.82+1.3/1.1/0+0.068 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 221 @34.083s 0%: 0.081+1.2+0.017 ms clock, 0.65+1.9/1.3/0+0.13 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 222 @34.086s 0%: 0.060+1.6+0.10 ms clock, 0.48+3.0/1.0/0+0.82 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 223 @34.091s 0%: 0.10+1.5+0.11 ms clock, 0.85+1.8/1.4/0+0.88 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 224 @34.094s 0%: 0.10+0.89+0.027 ms clock, 0.86+0.46/1.1/0+0.22 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 225 @34.096s 0%: 0.084+0.91+0.040 ms clock, 0.67+1.6/1.0/0+0.32 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 226 @34.099s 0%: 0.10+0.90+0.057 ms clock, 0.82+1.4/1.0/0+0.46 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 227 @34.102s 0%: 0.10+1.3+0.027 ms clock, 0.82+2.0/1.6/0+0.22 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 228 @34.106s 0%: 0.077+2.7+0.005 ms clock, 0.62+0.78/1.9/0+0.044 ms cpu, 6->8->3 MB, 7 MB goal, 8 P
gc 229 @34.112s 0%: 0.12+1.3+0.071 ms clock, 1.0+2.6/1.4/0+0.57 ms cpu, 7->7->3 MB, 8 MB goal, 8 P
gc 230 @34.115s 0%: 0.062+1.0+0.027 ms clock, 0.50+1.7/1.3/0+0.21 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 231 @34.118s 0%: 0.059+1.6+0.017 ms clock, 0.47+2.4/1.1/0+0.13 ms cpu, 5->7->3 MB, 6 MB goal, 8 P
gc 232 @34.121s 0%: 0.081+5.2+0.017 ms clock, 0.64+8.3/4.5/0+0.13 ms cpu, 5->8->5 MB, 7 MB goal, 8 P
gc 233 @34.130s 0%: 0.051+1.4+0.012 ms clock, 0.41+0.99/1.8/0+0.10 ms cpu, 8->9->3 MB, 10 MB goal, 8 P
gc 234 @34.132s 0%: 0.089+1.8+0.085 ms clock, 0.71+1.7/2.3/0+0.68 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 235 @34.136s 0%: 0.050+5.3+0.15 ms clock, 0.40+0.83/1.6/0+1.2 ms cpu, 6->9->5 MB, 7 MB goal, 8 P
gc 236 @34.144s 0%: 0.10+1.3+0.14 ms clock, 0.83+0.81/1.1/0+1.1 ms cpu, 9->9->3 MB, 11 MB goal, 8 P
gc 237 @34.147s 0%: 0.10+1.8+0.073 ms clock, 0.83+2.3/1.2/0+0.58 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 238 @34.150s 0%: 0.057+1.0+0.15 ms clock, 0.46+0.52/1.2/0+1.2 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 239 @34.153s 0%: 0.074+1.8+0.084 ms clock, 0.59+2.3/0.76/0+0.67 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 240 @34.156s 0%: 0.14+1.7+0.10 ms clock, 1.1+2.2/1.4/0+0.84 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 241 @34.159s 0%: 0.10+1.5+0.10 ms clock, 0.87+2.3/1.8/0+0.81 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 242 @34.163s 0%: 0.046+1.5+0.080 ms clock, 0.37+1.0/1.4/0+0.64 ms cpu, 5->7->3 MB, 6 MB goal, 8 P
gc 243 @34.166s 0%: 0.054+1.3+0.058 ms clock, 0.43+1.6/1.1/0+0.47 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 244 @34.168s 0%: 0.092+1.0+0.062 ms clock, 0.73+2.5/0.96/0+0.50 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 245 @34.170s 0%: 0.062+1.3+0.018 ms clock, 0.49+0.48/1.3/0+0.14 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 246 @34.173s 0%: 0.065+1.1+0.095 ms clock, 0.52+0.68/1.4/0+0.76 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 247 @34.176s 0%: 0.058+2.0+0.026 ms clock, 0.46+0.91/2.5/0+0.21 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 248 @34.180s 0%: 0.097+1.1+0.066 ms clock, 0.78+1.3/1.2/0+0.53 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 249 @34.183s 0%: 0.073+1.0+0.094 ms clock, 0.58+0.64/1.1/0+0.75 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 250 @34.185s 0%: 0.062+1.0+0.050 ms clock, 0.49+0.98/1.4/0+0.40 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 251 @34.188s 0%: 0.085+1.4+0.057 ms clock, 0.68+1.4/1.6/0+0.46 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 252 @34.191s 0%: 0.063+2.5+0.094 ms clock, 0.50+4.0/2.0/0+0.75 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 253 @34.196s 0%: 0.10+1.4+0.11 ms clock, 0.85+2.1/0.92/0+0.93 ms cpu, 7->8->3 MB, 8 MB goal, 8 P
gc 254 @34.199s 0%: 0.096+2.1+0.057 ms clock, 0.77+3.9/1.7/0+0.45 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 255 @34.202s 0%: 0.11+0.95+0.025 ms clock, 0.90+1.8/1.2/0+0.20 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 256 @34.204s 0%: 0.092+0.93+0.031 ms clock, 0.73+1.3/1.0/0+0.25 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 257 @34.206s 0%: 0.11+1.5+0.14 ms clock, 0.88+1.7/1.3/0+1.1 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 258 @34.209s 0%: 0.13+2.1+0.066 ms clock, 1.0+0.43/1.3/0+0.53 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 259 @34.213s 0%: 0.11+1.3+0.081 ms clock, 0.93+1.1/1.3/0+0.65 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 260 @34.216s 0%: 0.11+0.94+0.071 ms clock, 0.94+1.5/0.92/0+0.56 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 261 @34.218s 0%: 0.073+1.5+0.002 ms clock, 0.58+3.0/1.0/0+0.016 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 262 @34.221s 0%: 0.092+2.4+0.11 ms clock, 0.74+4.3/0/0+0.92 ms cpu, 5->7->3 MB, 6 MB goal, 8 P
gc 263 @34.224s 0%: 0.091+1.9+0.11 ms clock, 0.72+3.4/0.12/0+0.95 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 264 @34.228s 0%: 0.10+1.6+0.069 ms clock, 0.86+4.1/0.99/0+0.55 ms cpu, 5->6->3 MB, 7 MB goal, 8 P
gc 265 @34.230s 0%: 0.19+0.89+0.012 ms clock, 1.5+1.3/0.96/0+0.10 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 266 @34.232s 0%: 0.081+0.93+0.090 ms clock, 0.65+1.1/1.0/0+0.72 ms cpu, 4->5->3 MB, 5 MB goal, 8 P
gc 267 @34.234s 0%: 0.12+0.86+0.009 ms clock, 1.0+0.85/0.99/0+0.072 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 268 @34.236s 0%: 0.12+0.87+0.017 ms clock, 0.97+1.8/1.0/0+0.13 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 269 @34.238s 0%: 0.15+2.1+0.053 ms clock, 1.2+2.1/2.2/0+0.42 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 270 @34.242s 0%: 0.078+1.3+0.020 ms clock, 0.62+0.96/1.6/0+0.16 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 271 @34.245s 0%: 0.12+1.5+0.10 ms clock, 1.0+0.95/1.9/0+0.87 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 272 @34.248s 0%: 0.058+0.99+0.068 ms clock, 0.46+1.6/1.4/0+0.54 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 273 @34.250s 0%: 0.092+1.0+0.026 ms clock, 0.73+1.6/0.92/0+0.20 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 274 @34.253s 0%: 0.14+1.4+0.019 ms clock, 1.1+0.16/1.2/0+0.15 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 275 @34.256s 0%: 0.12+1.3+0.049 ms clock, 0.99+1.0/1.3/0+0.39 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 276 @34.259s 0%: 0.058+1.2+0.048 ms clock, 0.46+0.56/1.6/0+0.38 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 277 @34.262s 0%: 0.12+1.0+0.050 ms clock, 0.97+1.4/1.2/0+0.40 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 278 @34.265s 0%: 0.085+0.70+0.074 ms clock, 0.68+1.2/0.86/0+0.59 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 279 @34.267s 0%: 0.063+0.91+0.068 ms clock, 0.50+0.44/1.2/0+0.54 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 280 @34.269s 0%: 0.058+0.92+0.015 ms clock, 0.47+1.8/0.83/0+0.12 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 281 @34.271s 0%: 0.076+0.95+0.018 ms clock, 0.61+0.76/1.0/0+0.14 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 282 @34.274s 0%: 0.068+1.2+0.060 ms clock, 0.54+1.1/1.3/0+0.48 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 283 @34.277s 0%: 0.14+1.1+0.048 ms clock, 1.1+0.76/1.4/0+0.39 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 284 @34.281s 0%: 0.12+1.4+0.065 ms clock, 0.99+1.4/0.81/0+0.52 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 285 @34.284s 0%: 0.079+1.4+0.098 ms clock, 0.63+1.3/1.0/0+0.78 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 286 @34.287s 0%: 0.13+1.9+0.085 ms clock, 1.0+4.6/1.1/0+0.68 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 287 @34.291s 0%: 0.14+1.9+0.062 ms clock, 1.1+1.2/1.4/0+0.50 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 288 @34.295s 0%: 0.12+1.8+0.002 ms clock, 1.0+2.4/0.92/0+0.018 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 289 @34.298s 0%: 0.10+0.97+0.094 ms clock, 0.81+0.88/1.0/0+0.75 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 290 @34.301s 0%: 0.050+1.2+0.061 ms clock, 0.40+0.64/1.1/0+0.49 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 291 @34.303s 0%: 0.10+0.84+0.060 ms clock, 0.87+0.78/1.1/0+0.48 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 292 @34.306s 0%: 0.055+1.2+0.027 ms clock, 0.44+0.76/1.7/0+0.22 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 293 @34.309s 0%: 0.099+1.0+0.083 ms clock, 0.79+1.2/1.0/0+0.66 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 294 @34.312s 0%: 0.071+1.7+0.10 ms clock, 0.57+1.0/1.6/0+0.84 ms cpu, 5->7->3 MB, 6 MB goal, 8 P
gc 295 @34.315s 0%: 0.13+1.6+0.081 ms clock, 1.1+2.9/0.88/0+0.64 ms cpu, 7->7->3 MB, 8 MB goal, 8 P
gc 296 @34.320s 0%: 0.075+1.6+0.029 ms clock, 0.60+4.9/0/0+0.23 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 297 @34.323s 0%: 0.12+1.8+0.015 ms clock, 0.96+4.0/1.6/0+0.12 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 298 @34.326s 0%: 0.10+3.4+0.018 ms clock, 0.82+3.3/1.4/0+0.14 ms cpu, 5->7->4 MB, 6 MB goal, 8 P
gc 299 @34.331s 0%: 0.076+1.3+0.12 ms clock, 0.60+0.74/0.91/0+1.0 ms cpu, 6->7->3 MB, 8 MB goal, 8 P
gc 300 @34.334s 0%: 0.13+1.0+0.12 ms clock, 1.1+1.6/1.1/0+0.97 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 301 @34.336s 0%: 0.18+1.2+0.078 ms clock, 1.4+1.8/1.6/0+0.62 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 302 @34.339s 0%: 0.058+2.1+0.014 ms clock, 0.47+5.2/1.8/0+0.11 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 303 @34.343s 0%: 0.078+1.7+0.018 ms clock, 0.62+0.58/2.2/0+0.15 ms cpu, 5->6->3 MB, 7 MB goal, 8 P
gc 304 @34.346s 0%: 0.11+1.6+0.079 ms clock, 0.90+1.1/1.3/0+0.63 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 305 @34.349s 0%: 0.084+0.83+0.067 ms clock, 0.67+1.6/0.87/0+0.53 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 306 @34.351s 0%: 0.093+1.6+0.026 ms clock, 0.74+1.3/2.3/0+0.21 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 307 @34.354s 0%: 0.050+1.1+0.12 ms clock, 0.40+0.86/1.4/0+1.0 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 308 @34.357s 0%: 0.13+1.4+0.12 ms clock, 1.0+0.92/1.4/0+1.0 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 309 @34.360s 0%: 0.069+1.8+0.11 ms clock, 0.55+0.14/1.3/0+0.92 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 310 @34.364s 0%: 0.079+1.0+0.017 ms clock, 0.63+0.74/1.2/0+0.14 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 311 @34.366s 0%: 0.10+1.0+0.014 ms clock, 0.86+1.1/1.1/0+0.11 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 312 @34.368s 0%: 0.094+1.1+0.13 ms clock, 0.75+1.6/1.2/0+1.0 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 313 @34.371s 0%: 0.043+1.1+0.088 ms clock, 0.35+0.83/1.2/0+0.70 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 314 @34.374s 0%: 0.12+1.2+0.10 ms clock, 0.96+1.5/1.3/0+0.83 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 315 @34.378s 0%: 0.084+1.2+0.063 ms clock, 0.67+0.22/1.0/0+0.50 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 316 @34.381s 0%: 0.096+1.0+0.016 ms clock, 0.76+0.89/1.5/0+0.13 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 317 @34.385s 0%: 0.088+3.3+0.018 ms clock, 0.70+1.7/1.4/0+0.14 ms cpu, 6->7->4 MB, 7 MB goal, 8 P
gc 318 @34.393s 0%: 0.11+2.9+0.029 ms clock, 0.92+2.4/1.6/0+0.23 ms cpu, 7->8->3 MB, 8 MB goal, 8 P
gc 319 @34.400s 0%: 0.079+3.3+0.020 ms clock, 0.63+1.6/3.9/0+0.16 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 320 @34.406s 0%: 0.11+3.7+0.018 ms clock, 0.93+4.0/3.4/0+0.14 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 321 @34.416s 0%: 0.093+3.7+0.015 ms clock, 0.74+1.1/1.8/0+0.12 ms cpu, 5->7->3 MB, 6 MB goal, 8 P
gc 322 @34.424s 0%: 0.067+5.0+0.019 ms clock, 0.53+5.6/5.5/0.001+0.15 ms cpu, 6->8->4 MB, 7 MB goal, 8 P
gc 323 @34.433s 0%: 0.18+3.3+0.012 ms clock, 1.4+1.9/2.9/0+0.10 ms cpu, 6->8->3 MB, 8 MB goal, 8 P
gc 324 @34.440s 0%: 0.11+3.0+0.10 ms clock, 0.92+0.67/2.7/0+0.80 ms cpu, 6->8->3 MB, 7 MB goal, 8 P
gc 325 @34.445s 0%: 0.064+1.3+0.027 ms clock, 0.51+0.26/1.0/0+0.21 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 326 @34.448s 0%: 0.14+1.3+0.056 ms clock, 1.1+0.10/1.6/0+0.45 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 327 @34.451s 0%: 0.099+1.8+0.061 ms clock, 0.79+2.3/2.2/0+0.49 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 328 @34.457s 0%: 0.56+2.1+0.013 ms clock, 4.5+0.94/2.0/0+0.10 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 329 @34.462s 0%: 0.12+1.5+0.037 ms clock, 0.96+0.80/1.0/0+0.30 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 330 @34.466s 0%: 0.12+0.88+0.039 ms clock, 1.0+0.70/1.1/0+0.31 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 331 @34.469s 0%: 0.21+1.4+0.054 ms clock, 1.7+0.52/0.86/0+0.43 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 332 @34.473s 0%: 0.37+3.9+0.018 ms clock, 3.0+0.94/1.4/0+0.14 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 333 @34.481s 0%: 0.14+1.6+0.060 ms clock, 1.1+0.16/1.5/0+0.48 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 334 @34.486s 0%: 0.074+1.4+0.012 ms clock, 0.59+0.67/1.1/0+0.10 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 335 @34.491s 0%: 0.25+1.4+0.047 ms clock, 2.0+0.88/1.9/0+0.38 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 336 @34.494s 0%: 0.25+2.3+0.020 ms clock, 2.0+1.0/1.2/0+0.16 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 337 @34.500s 0%: 0.061+1.6+0.011 ms clock, 0.48+0.49/1.4/0+0.089 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 338 @34.504s 0%: 0.13+1.6+0.075 ms clock, 1.1+1.5/1.7/0+0.60 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 339 @34.508s 0%: 0.13+1.6+0.077 ms clock, 1.0+0.40/1.6/0+0.61 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 340 @34.512s 0%: 0.23+1.4+0.15 ms clock, 1.8+1.5/1.7/0+1.2 ms cpu, 6->6->3 MB, 7 MB goal, 8 P
gc 341 @34.517s 0%: 0.17+1.6+0.090 ms clock, 1.3+1.6/1.7/0+0.72 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 342 @34.521s 0%: 0.13+1.8+0.078 ms clock, 1.0+1.6/1.7/0+0.63 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 343 @34.526s 0%: 0.096+3.1+0.023 ms clock, 0.77+0.94/1.8/0+0.18 ms cpu, 6->7->4 MB, 7 MB goal, 8 P
gc 344 @34.532s 0%: 0.11+1.5+0.11 ms clock, 0.88+2.1/1.4/0+0.90 ms cpu, 7->7->3 MB, 8 MB goal, 8 P
gc 345 @34.536s 0%: 0.071+1.1+0.024 ms clock, 0.57+1.0/1.4/0+0.19 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 346 @34.539s 0%: 0.16+1.5+0.002 ms clock, 1.3+1.0/1.4/0+0.022 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 347 @34.543s 0%: 0.083+1.9+0.002 ms clock, 0.66+3.9/0.92/0+0.020 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 348 @34.547s 0%: 0.070+2.0+0.072 ms clock, 0.56+0.77/1.5/0+0.57 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 349 @34.552s 0%: 0.11+2.7+0.066 ms clock, 0.89+0.73/2.8/0+0.53 ms cpu, 7->9->4 MB, 8 MB goal, 8 P
gc 350 @34.558s 0%: 0.21+1.2+0.051 ms clock, 1.7+1.1/1.6/0+0.40 ms cpu, 7->8->2 MB, 8 MB goal, 8 P
gc 351 @34.562s 0%: 0.076+1.3+0.090 ms clock, 0.61+1.7/1.6/0+0.72 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 352 @34.565s 0%: 0.093+0.98+0.026 ms clock, 0.74+1.0/1.3/0+0.21 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 353 @34.568s 0%: 0.068+2.6+0.008 ms clock, 0.54+2.7/1.6/0+0.065 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 354 @34.575s 0%: 0.10+4.9+0.091 ms clock, 0.81+1.0/5.0/0+0.72 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 355 @34.584s 0%: 0.13+4.2+0.018 ms clock, 1.0+2.0/3.5/0+0.14 ms cpu, 5->7->3 MB, 6 MB goal, 8 P
gc 356 @34.592s 0%: 0.14+2.7+0.014 ms clock, 1.1+4.1/2.3/0+0.11 ms cpu, 6->8->3 MB, 7 MB goal, 8 P
gc 357 @34.596s 0%: 0.10+1.9+0.025 ms clock, 0.80+1.5/2.5/0+0.20 ms cpu, 6->7->3 MB, 7 MB goal, 8 P
gc 358 @34.600s 0%: 0.064+2.5+0.020 ms clock, 0.51+0.98/1.5/0+0.16 ms cpu, 5->7->3 MB, 6 MB goal, 8 P
gc 359 @34.605s 0%: 0.074+3.4+0.004 ms clock, 0.59+1.8/4.4/0.082+0.036 ms cpu, 5->8->4 MB, 7 MB goal, 8 P
gc 360 @34.612s 0%: 0.079+1.1+0.088 ms clock, 0.63+1.1/1.1/0+0.71 ms cpu, 6->7->2 MB, 8 MB goal, 8 P
gc 361 @34.615s 0%: 0.068+0.73+0.13 ms clock, 0.54+0.069/0.91/0+1.0 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
gc 362 @34.617s 0%: 0.083+1.0+0.073 ms clock, 0.67+0.65/0.93/0+0.58 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 363 @34.619s 0%: 0.092+2.4+0.076 ms clock, 0.73+1.4/3.9/0+0.60 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 364 @34.624s 0%: 0.15+3.5+0.12 ms clock, 1.2+2.4/1.4/0+1.0 ms cpu, 6->8->3 MB, 7 MB goal, 8 P
gc 365 @34.630s 0%: 0.069+1.0+0.065 ms clock, 0.55+0.14/0.93/0+0.52 ms cpu, 6->7->2 MB, 7 MB goal, 8 P
gc 366 @34.633s 0%: 0.087+1.2+0.13 ms clock, 0.70+0.12/1.5/0+1.1 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 367 @34.635s 0%: 0.097+2.8+0.018 ms clock, 0.77+0.73/3.2/0+0.14 ms cpu, 5->7->4 MB, 6 MB goal, 8 P
gc 368 @34.645s 0%: 0.10+0.89+0.013 ms clock, 0.87+0.10/0.83/0+0.10 ms cpu, 8->8->2 MB, 9 MB goal, 8 P
gc 369 @34.647s 0%: 0.15+1.2+0.019 ms clock, 1.2+1.8/1.6/0+0.15 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 370 @34.651s 0%: 0.067+1.2+0.14 ms clock, 0.53+0.12/1.0/0+1.1 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 371 @34.654s 0%: 0.14+4.3+0.069 ms clock, 1.1+1.3/1.5/0+0.55 ms cpu, 5->7->4 MB, 6 MB goal, 8 P
gc 372 @34.662s 0%: 0.18+2.1+0.16 ms clock, 1.4+2.3/1.7/0+1.3 ms cpu, 7->8->3 MB, 8 MB goal, 8 P
gc 373 @34.666s 0%: 0.083+3.3+0.032 ms clock, 0.67+2.5/3.2/0+0.25 ms cpu, 6->8->4 MB, 7 MB goal, 8 P
gc 374 @34.671s 0%: 0.12+0.85+0.073 ms clock, 1.0+1.6/0.97/0+0.59 ms cpu, 6->7->2 MB, 8 MB goal, 8 P
gc 375 @34.673s 0%: 0.15+2.3+0.012 ms clock, 1.2+0.93/1.2/0+0.096 ms cpu, 4->5->3 MB, 5 MB goal, 8 P
gc 376 @34.678s 0%: 0.053+5.4+0.014 ms clock, 0.42+0.79/10/0.23+0.11 ms cpu, 5->7->4 MB, 7 MB goal, 8 P
gc 377 @34.685s 0%: 0.060+3.4+0.050 ms clock, 0.48+0.99/3.6/0+0.40 ms cpu, 6->8->4 MB, 8 MB goal, 8 P
gc 378 @34.691s 0%: 0.069+1.2+0.020 ms clock, 0.55+1.0/1.9/1.2+0.16 ms cpu, 7->7->2 MB, 9 MB goal, 8 P
gc 379 @34.696s 0%: 0.046+0.73+0.021 ms clock, 0.37+0.67/1.2/0.31+0.16 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
GC forced
gc 380 @154.760s 0%: 0.30+0.51+0.002 ms clock, 2.4+0/0.91/1.2+0.016 ms cpu, 3->3->2 MB, 4 MB goal, 8 P
```

and garbage colelction has been forced 


Here's how one of these lines breaks down. STW = stop the world:
gc 380                                  @154.760s                                0%:
gc nth GC run since program started     time elapsed since program started       percent of available CPU so far spent on GC

0.30+0.51+0.002 ms clock
wall clock. broken down further:
0.30  - STW : write barrier : wait for all pointers to reach a GC-safe point.
0.51  - concurrent: marking
0.002 - STW: mark term - write barrier off and clean up.

2.4+0/0.91/1.2+0.016 ms cpu
CPU time
2.4   - STW: write barrier
0.91  - concurrent: mark - assist time
1.2   - concurrent: mark - idle GC time
0.016 - STW: mark term

3->3->2 MB 
breaks down to
3 MB heap memory in use before marking started
3 MB heap memory in use after marking finished
2 MB heap memory marked as live after marking

4 MB goal 
collection goal for heap memory in-use after marking finished


8 P
number of logical processes or threads used to run goroutines


None of the truns seem particularly expensive on inspection, there's just a LOT (380) of them.

The pprof profile is pretty helpful, especially when visualized as a png. We're using >600 MB allocating strings, out of 790MB total.

![Initial allocation graph](./profile001.svg)


Initial run of the benchmark test:

```
❯ go test -bench=.
2022/09/22 11:40:43 reloaded cache http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml
goos: darwin
goarch: arm64
pkg: github.com/ardanlabs/gotraining/topics/go/profiling/project/search
BenchmarkRssSearch-8       79545             14490 ns/op
PASS
ok      github.com/ardanlabs/gotraining/topics/go/profiling/project/search      2.237s
```

Ok! Much improved after getting rid of a bunch of string allocation.

```
❯ go build && GODEBUG=gctrace=1 ./project
2022/09/22 11:52:11.305346 service.go:64: Listening on: 0.0.0.0:5000
gc 1 @3.494s 0%: 0.018+0.52+0.001 ms clock, 0.15+0.22/0.91/0.46+0.012 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
2022/09/22 11:52:14.815241 rss.go:108: reloaded cache http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml
gc 2 @3.526s 0%: 0.031+0.82+0.002 ms clock, 0.25+0.41/1.3/1.8+0.017 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
2022/09/22 11:52:14.840269 rss.go:108: reloaded cache http://rss.nytimes.com/services/xml/rss/nyt/US.xml
2022/09/22 11:52:14.861222 rss.go:108: reloaded cache http://rss.nytimes.com/services/xml/rss/nyt/Politics.xml
2022/09/22 11:52:14.866597 rss.go:108: reloaded cache http://rss.cnn.com/rss/cnn_topstories.rss
gc 3 @3.564s 0%: 0.036+1.0+0.060 ms clock, 0.29+0.66/1.2/1.9+0.48 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
2022/09/22 11:52:14.886677 rss.go:108: reloaded cache http://rss.nytimes.com/services/xml/rss/nyt/Business.xml
2022/09/22 11:52:14.909880 rss.go:108: reloaded cache http://feeds.bbci.co.uk/news/rss.xml
2022/09/22 11:52:14.988839 rss.go:108: reloaded cache http://feeds.bbci.co.uk/news/world/rss.xml
2022/09/22 11:52:14.999826 rss.go:108: reloaded cache http://rss.cnn.com/rss/cnn_world.rss
2022/09/22 11:52:15.074179 rss.go:108: reloaded cache http://feeds.bbci.co.uk/news/politics/rss.xml
gc 4 @3.772s 0%: 0.17+1.2+0.14 ms clock, 1.4+0.73/1.9/2.9+1.1 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
2022/09/22 11:52:15.148478 rss.go:108: reloaded cache http://rss.cnn.com/rss/cnn_us.rss
2022/09/22 11:52:15.168065 rss.go:108: reloaded cache http://feeds.bbci.co.uk/news/world/us_and_canada/rss.xml
2022/09/22 11:52:15.293728 rss.go:108: reloaded cache http://rss.cnn.com/rss/cnn_allpolitics.rss
gc 5 @3.993s 0%: 0.32+2.2+0.11 ms clock, 2.6+2.2/2.3/0+0.93 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 6 @4.001s 0%: 0.66+1.2+0.13 ms clock, 5.3+2.6/1.7/0+1.0 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 7 @4.009s 0%: 0.51+0.98+0.13 ms clock, 4.1+2.2/1.3/0+1.0 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
gc 8 @4.016s 0%: 0.26+1.8+0.020 ms clock, 2.1+1.0/1.7/0+0.16 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 9 @4.025s 0%: 0.18+1.0+0.066 ms clock, 1.5+1.8/1.2/0+0.52 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 10 @4.032s 0%: 0.22+2.0+0.025 ms clock, 1.7+2.1/2.0/0+0.20 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 11 @4.039s 0%: 0.21+2.3+0.11 ms clock, 1.6+1.6/1.7/0+0.95 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 12 @4.048s 0%: 0.33+0.92+0.083 ms clock, 2.6+1.3/1.2/0+0.66 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 13 @4.054s 0%: 0.064+1.2+0.10 ms clock, 0.51+1.4/1.5/0+0.80 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 14 @4.060s 0%: 0.34+1.1+0.077 ms clock, 2.7+1.3/1.2/0+0.61 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 15 @4.066s 0%: 0.14+0.76+0.039 ms clock, 1.1+1.3/0.92/0+0.31 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 16 @4.070s 0%: 0.24+1.0+0.004 ms clock, 1.9+1.2/0.88/0+0.035 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
gc 17 @4.076s 0%: 0.086+0.63+0.032 ms clock, 0.69+0.75/0.79/0+0.26 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 18 @4.080s 0%: 0.13+0.83+0.049 ms clock, 1.1+1.5/0.91/0+0.39 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 19 @4.086s 0%: 0.11+0.88+0.023 ms clock, 0.91+1.1/0.91/0+0.19 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 20 @4.091s 0%: 0.15+0.78+0.076 ms clock, 1.2+1.8/0.88/0+0.61 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 21 @4.095s 0%: 0.11+1.0+0.004 ms clock, 0.94+1.8/1.1/0+0.032 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 22 @4.101s 0%: 0.13+0.91+0.012 ms clock, 1.0+1.4/0.84/0+0.10 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 23 @4.106s 0%: 0.16+0.89+0.017 ms clock, 1.3+0.58/0.87/0+0.14 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 24 @4.111s 0%: 0.11+0.81+0.080 ms clock, 0.93+1.5/1.0/0+0.64 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 25 @4.116s 0%: 0.093+1.6+0.046 ms clock, 0.74+0.68/0.89/0+0.37 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 26 @4.123s 0%: 0.11+0.82+0.045 ms clock, 0.90+1.0/1.0/0+0.36 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 27 @4.128s 0%: 0.088+2.6+0.11 ms clock, 0.70+0.81/0.86/0+0.88 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 28 @4.137s 0%: 0.21+1.2+0.10 ms clock, 1.7+1.4/1.0/0+0.86 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 29 @4.143s 0%: 0.17+1.0+0.018 ms clock, 1.3+1.7/1.0/0+0.14 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 30 @4.149s 0%: 0.11+2.1+0.082 ms clock, 0.94+1.4/0.79/0+0.65 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 31 @4.157s 0%: 0.18+0.65+0.043 ms clock, 1.5+0.35/0.67/0+0.35 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 32 @4.161s 0%: 0.084+0.79+0.038 ms clock, 0.67+0.62/0.87/0+0.30 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
gc 33 @4.166s 0%: 0.11+0.89+0.11 ms clock, 0.89+1.3/0.82/0+0.91 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 34 @4.171s 0%: 0.092+1.4+0.091 ms clock, 0.73+0.57/0.88/0+0.72 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 35 @4.178s 0%: 0.093+0.81+0.022 ms clock, 0.74+1.5/1.0/0+0.17 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 36 @4.183s 0%: 0.14+0.96+0.042 ms clock, 1.1+0.52/1.0/0+0.33 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 37 @4.188s 0%: 0.14+0.81+0.10 ms clock, 1.1+1.2/0.82/0+0.80 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 38 @4.193s 0%: 0.16+1.0+0.045 ms clock, 1.2+1.2/0.90/0+0.36 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 39 @4.198s 0%: 0.082+0.70+0.076 ms clock, 0.65+1.3/0.69/0+0.60 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 40 @4.203s 0%: 0.18+0.83+0.062 ms clock, 1.5+0.66/0.90/0+0.49 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
gc 41 @4.207s 0%: 0.12+0.85+0.056 ms clock, 1.0+1.1/0.83/0+0.45 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 42 @4.212s 0%: 0.096+1.3+0.020 ms clock, 0.77+0.79/1.2/0+0.16 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 43 @4.218s 0%: 0.12+1.0+0.061 ms clock, 0.98+1.3/0.83/0+0.48 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 44 @4.224s 0%: 0.062+0.82+0.095 ms clock, 0.50+0.13/0.93/0+0.76 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 45 @4.229s 0%: 0.12+0.99+0.10 ms clock, 1.0+1.3/1.0/0+0.81 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 46 @4.234s 0%: 0.15+1.0+0.073 ms clock, 1.2+1.2/0.78/0+0.59 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 47 @4.239s 0%: 0.12+0.70+0.10 ms clock, 0.99+1.2/0.76/0+0.85 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 48 @4.244s 0%: 0.10+0.79+0.059 ms clock, 0.84+1.6/0.95/0+0.47 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 49 @4.249s 0%: 0.033+1.0+0.014 ms clock, 0.26+0.24/0.87/0+0.11 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 50 @4.254s 0%: 0.10+0.73+0.017 ms clock, 0.85+0.62/0.95/0+0.13 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 51 @4.259s 0%: 0.095+0.64+0.014 ms clock, 0.76+0.63/0.74/0+0.11 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 52 @4.264s 0%: 0.17+0.87+0.024 ms clock, 1.4+0.42/1.0/0+0.19 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 53 @4.270s 0%: 0.14+1.2+0.019 ms clock, 1.1+1.0/1.3/0+0.15 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 54 @4.276s 0%: 0.18+1.5+0.029 ms clock, 1.4+1.4/0.87/0+0.23 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 55 @4.282s 0%: 0.059+0.78+0.030 ms clock, 0.47+0.24/0.73/0+0.24 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 56 @4.287s 0%: 0.15+0.81+0.028 ms clock, 1.2+0.92/0.87/0+0.22 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 57 @4.292s 0%: 0.081+0.92+0.081 ms clock, 0.65+1.4/1.0/0+0.65 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 58 @4.297s 0%: 0.094+1.0+0.017 ms clock, 0.75+0.31/0.83/0+0.13 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 59 @4.303s 0%: 0.14+1.0+0.016 ms clock, 1.1+1.4/1.1/0+0.13 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 60 @4.310s 0%: 0.12+0.79+0.079 ms clock, 0.98+1.3/0.98/0+0.63 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 61 @4.315s 0%: 0.060+3.1+0.036 ms clock, 0.48+0.56/3.6/0+0.29 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
gc 62 @4.326s 0%: 0.061+0.88+0.015 ms clock, 0.49+0.67/1.0/0+0.12 ms cpu, 6->6->2 MB, 7 MB goal, 8 P
gc 63 @4.331s 0%: 0.052+0.84+0.056 ms clock, 0.41+0.25/0.82/0+0.45 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 64 @4.338s 0%: 0.058+0.81+0.016 ms clock, 0.47+0.37/0.74/0+0.13 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 65 @4.344s 0%: 0.086+1.0+0.027 ms clock, 0.69+1.6/1.2/0+0.21 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 66 @4.351s 0%: 0.14+0.77+0.017 ms clock, 1.1+1.5/0.73/0+0.14 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 67 @4.357s 0%: 0.034+0.84+0.001 ms clock, 0.27+0.67/0.92/0+0.012 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 68 @4.363s 0%: 0.040+1.0+0.067 ms clock, 0.32+0.68/1.1/0+0.53 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 69 @4.369s 0%: 0.051+0.81+0.054 ms clock, 0.40+0.94/0.90/0+0.43 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 70 @4.377s 0%: 0.040+2.2+0.045 ms clock, 0.32+2.6/1.9/0+0.36 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 71 @4.384s 0%: 0.044+0.75+0.087 ms clock, 0.35+0.63/0.74/0+0.70 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 72 @4.390s 0%: 0.045+2.1+0.022 ms clock, 0.36+0.33/2.4/0+0.18 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 73 @4.400s 0%: 0.051+0.89+0.062 ms clock, 0.40+0.91/0.92/0+0.50 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 74 @4.406s 0%: 0.051+0.90+0.021 ms clock, 0.41+0.42/0.68/0+0.17 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 75 @4.414s 0%: 0.044+0.72+0.035 ms clock, 0.35+1.5/0.75/0+0.28 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 76 @4.421s 0%: 0.14+1.0+0.036 ms clock, 1.1+1.4/1.2/0+0.29 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 77 @4.429s 0%: 0.068+1.6+0.009 ms clock, 0.54+1.0/0.99/0+0.073 ms cpu, 5->6->3 MB, 6 MB goal, 8 P
gc 78 @4.438s 0%: 0.062+1.5+0.083 ms clock, 0.49+0.80/1.0/0+0.66 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
gc 79 @4.446s 0%: 0.054+6.6+0.038 ms clock, 0.43+0.38/0.85/0+0.30 ms cpu, 5->7->4 MB, 6 MB goal, 8 P
gc 80 @4.464s 0%: 0.10+0.86+0.050 ms clock, 0.84+1.3/1.1/0+0.40 ms cpu, 8->8->2 MB, 9 MB goal, 8 P
gc 81 @4.470s 0%: 0.058+0.76+0.068 ms clock, 0.47+0.36/0.76/0+0.54 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 82 @4.476s 0%: 0.098+0.61+0.057 ms clock, 0.78+0.92/0.72/0+0.45 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 83 @4.480s 0%: 0.20+1.0+0.090 ms clock, 1.6+0.11/1.1/0+0.72 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
gc 84 @4.485s 0%: 0.053+0.81+0.073 ms clock, 0.42+0.51/0.90/0+0.59 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 85 @4.491s 0%: 0.086+0.92+0.022 ms clock, 0.69+0.32/0.76/0+0.18 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 86 @4.498s 0%: 0.028+1.0+0.060 ms clock, 0.22+0.77/0.88/0+0.48 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 87 @4.503s 0%: 0.10+0.72+0.013 ms clock, 0.82+0.36/0.71/0+0.10 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 88 @4.509s 0%: 0.045+1.5+0.002 ms clock, 0.36+0.43/1.6/0+0.017 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 89 @4.517s 0%: 0.038+0.86+0.019 ms clock, 0.30+0.28/0.54/0+0.15 ms cpu, 5->5->2 MB, 6 MB goal, 8 P
gc 90 @4.524s 0%: 0.10+0.67+0.014 ms clock, 0.83+0.34/0.63/0+0.11 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
gc 91 @4.529s 0%: 0.039+0.66+0.001 ms clock, 0.31+0.26/0.72/0+0.009 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
gc 92 @4.534s 0%: 0.036+1.0+0.016 ms clock, 0.29+1.1/0.65/0+0.13 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 93 @4.540s 0%: 0.083+0.53+0.036 ms clock, 0.66+0.39/0.61/0+0.29 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 94 @4.544s 0%: 0.042+0.39+0.006 ms clock, 0.34+0.19/0.61/0.064+0.049 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
```

That's about a quarter of the GC runs. Here's the graph.
![Allocation graph after cleaning up strings](./profile002.svg)

