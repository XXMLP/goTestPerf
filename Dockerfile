FROM harbor.dx-corp.top/basic/python:3-stretch

RUN pip install --no-cache-dir -i https://pypi.douban.com/simple git+https://github.com/feiyuw/locust.git@sstore
RUN mkdir /work

COPY dummy.py /work

EXPOSE 8089 5557 5558

ENTRYPOINT ["locust", "-f", "/work/dummy.py", "--master"]
