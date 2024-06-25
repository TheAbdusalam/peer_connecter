FROM centos:latest

COPY ./ /app

# Add mirror lists
RUN curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo && \
    curl -o /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-7.repo && \
    yum -y install epel-release && \
    yum -y install git go && \
    pip3 install -r /app/requirements.txt && \
    rm -rf /var/cache/yum
WORKDIR /app

ENTRYPOINT ["/bin/sh"]
