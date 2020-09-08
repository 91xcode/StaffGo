<?php
// GENERATED CODE -- DO NOT EDIT!

namespace User;

/**
 */
class UserClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \User\UserIndexRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \User\UserIndexResponse
     */
    public function UserIndex(\User\UserIndexRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/user.User/UserIndex',
        $argument,
        ['\User\UserIndexResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \User\UserViewRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \User\UserViewResponse
     */
    public function UserView(\User\UserViewRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/user.User/UserView',
        $argument,
        ['\User\UserViewResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \User\UserPostRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \User\UserPostResponse
     */
    public function UserPost(\User\UserPostRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/user.User/UserPost',
        $argument,
        ['\User\UserPostResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \User\UserDeleteRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \User\UserDeleteResponse
     */
    public function UserDelete(\User\UserDeleteRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/user.User/UserDelete',
        $argument,
        ['\User\UserDeleteResponse', 'decode'],
        $metadata, $options);
    }

}
